package storageserver

import (
	"path"
	"archive/zip"
	"strings"
	"path/filepath"
	"fmt"
	"os"
	"io"
	fileutil "theterriblechild/CloudApp/tools/utils/file"
)

const (
	CACHE_FOLDER = "Cache"
)

func StoreToUserCache(cacheLocation string, userID string, agentID string, src string, key []byte) error {
	userCacheLocation := GetUserAgentCacheLocation(cacheLocation, userID, agentID)
	_, err := UnzipCompressEncrypt(src, userCacheLocation, key)
	return err
}

func FromUserCache(cacheLocation string, userID string, agentID string, files []string, dest string, key []byte) error {
	userCacheLocation := GetUserAgentCacheLocation(cacheLocation, userID, agentID)
	for i, _ := range files {
		files[i] = path.Join(userCacheLocation, files[i])
	}
	return DecryptDecompressZip(files, dest, key, false)
}

func GetUserAgentCacheLocation(cacheLocation string, userID string, agentID string) string {
	return path.Join(cacheLocation, userID, CACHE_FOLDER, agentID)
}

func DecryptDecompressZip(files []string, dest string, key []byte, useAbsolutePath bool) error {
	if len(dest) == 0 {
		return fmt.Errorf("No output file path given.")
	}

	var removeIdx int
	if !useAbsolutePath {
		commonPrefix := fileutil.GetCommonPrefix(files)
		removeIdx = len(commonPrefix)
	}

	newZipFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {

		inFile, err := os.Open(file)
		if err != nil {
			return err
		}
		info, err := inFile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = file[removeIdx:]
		header.Method = zip.Deflate

		reader, err := fileutil.GetGZipAESReader(inFile, key)
		if err != nil {
			return err
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		for {
			if _, err := io.Copy(writer, reader); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
		}
		reader.Close()
	}
	return nil
}

func UnzipCompressEncrypt(src string, dest string, key []byte) ([]string, error) {
	var filenames []string

    r, err := zip.OpenReader(src)
    if err != nil {
        return filenames, err
    }
    defer r.Close()

    for _, f := range r.File {

        rc, err := f.Open()
        if err != nil {
            return filenames, err
        }
        defer rc.Close()

        fpath := filepath.Join(dest, f.Name)

        if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
            return filenames, fmt.Errorf("%s: illegal file path", fpath)
        }

        filenames = append(filenames, fpath)

        if f.FileInfo().IsDir() {
            os.MkdirAll(fpath, os.ModePerm)
        } else {
            if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
                return filenames, err
            }

            outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return filenames, err
			}
			
			writer, err := fileutil.GetGZipAESWriter(outFile, key)
			if err != nil {
				return filenames, err
			}

            _, err = io.Copy(writer, rc)
			writer.Close()
            outFile.Close()

            if err != nil {
                return filenames, err
            }

        }
    }
    return filenames, nil
}