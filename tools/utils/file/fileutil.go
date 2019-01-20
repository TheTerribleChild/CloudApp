package fileutil

import (
	"archive/zip"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"theterriblechild/CloudApp/tools/utils/hash"
)

type MergeMode uint8
const (
	BlockSize int64 = 1024 * 1024
	OVERWRITE MergeMode = 0
	RENAME    MergeMode = 1
	SKIP      MergeMode = 2
)

//GetAllFileInDirectoryRecursively Gets all regular file with extension in root directory recursively
func GetAllFileInDirectoryRecursively(roots []string, ext string) ([]string, error) {
	fileList := []string{}
	var err error
	for _, root := range roots {
		err = filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if f.Mode().IsRegular() {
				if (len(ext) != 0 && filepath.Ext(path) == ext) || len(ext) == 0 {
					fileList = append(fileList, path)
				}
			}
			return nil
		})
	}
	return fileList, err
}

//GetAllFileInDirectory Gets all regular file with extension in root directory
func GetAllFileInDirectory(root string, ext string) ([]string, error) {
	fileList := []string{}
	files, err := ioutil.ReadDir(root)
	for _, f := range files {
		if (len(ext) != 0 && filepath.Ext(f.Name()) == ext) || len(ext) == 0 {
			fileList = append(fileList, f.Name())
		}
	}
	return fileList, err
}

//Get size of file/dir recursively
func GetFileSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func ZipFiles(files []string, dest string, useAbsolutePath bool) error {
	if len(dest) == 0 {
		return fmt.Errorf("No output file path given.")
	}

	var removeIdx int = 0
	if !useAbsolutePath {
		commonPrefix := GetCommonPrefix(files)
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

		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = file[removeIdx:]
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		for {
			if _, err := io.CopyN(writer, zipfile, BlockSize); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
		}
		zipfile.Close()
	}
	return nil
}

func Unzip(src string, dest string) ([]string, error) {

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

            _, err = io.Copy(outFile, rc)

            outFile.Close()

            if err != nil {
                return filenames, err
            }

        }
    }
    return filenames, nil
}

func GetCommonPrefix(paths []string) string {
	// Handle special cases.
	sep := byte(os.PathSeparator)
	switch len(paths) {
	case 0:
		return ""
	case 1:
		return path.Clean(paths[0])
	}

	c := []byte(path.Clean(paths[0]))
	c = append(c, sep)

	// Ignore the first path since it's already in c
	for _, v := range paths[1:] {
		// Clean up each path before testing it
		v = path.Clean(v) + string(sep)

		// Find the first non-common byte and truncate c
		if len(v) < len(c) {
			c = c[:len(v)]
		}
		for i := 0; i < len(c); i++ {
			if v[i] != c[i] {
				c = c[:i]
				break
			}
		}
	}

	// Remove trailing non-separator characters and the final separator
	for i := len(c) - 1; i >= 0; i-- {
		if c[i] == sep {
			c = c[:i]
			break
		}
	}

	return string(c)
}

func EncryptFile(src string, dest string, key []byte) error {
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}
	inFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cipherWriter, err := GetAESEncryptionWriter(outFile, key)
	if _, err := io.Copy(cipherWriter, inFile); err != nil {
		return err
	}
	os.Chmod(dest, stat.Mode())
	return nil
}

func DecryptFile(src string, dest string, key []byte) error {
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}
	inFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cipherReader, err := GetAESEncryptionReader(inFile, key)
	if _, err := io.Copy(outFile, cipherReader); err != nil {
		return err
	}
	os.Chmod(dest, stat.Mode())
	return nil
}

// func CompressFile(src string, dest string) error {
// 	if _, err := os.Stat(src); err != nil {
// 		return err
// 	}
// }

// func DecompressFile(src string, dest string) error {
// 	if _, err := os.Stat(src); err != nil {
// 		return err
// 	}
// }

func CompressAndEncryptFile(src string, dest string, key []byte) error {
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}
	inFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inFile.Close()
	outFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer, err := GetGZipAESWriter(outFile, key)
	if err != nil {
		return err
	}
	defer writer.Close()
	if _, err := io.Copy(writer, inFile); err != nil {
		return err
	}

	os.Chmod(dest, stat.Mode())
	return nil
}

func DecryptAndDecompressFile(src string, dest string, key []byte) error {
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}
	inFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inFile.Close()
	outFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer outFile.Close()

	reader, err := GetGZipAESReader(inFile, key)
	if err != nil {
		return err
	}
	reader.Close()
	if _, err := io.Copy(outFile, reader); err != nil {
		return err
	}

	os.Chmod(dest, stat.Mode())
	return nil
}

func GetAESEncryptionWriter(writer io.Writer, key []byte) (*cipher.StreamWriter, error){
	block, err := aes.NewCipher(key)
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])
	cipherWriter := &cipher.StreamWriter{S: stream, W: writer}
	return cipherWriter, err
}

func GetAESEncryptionReader(reader io.Reader, key []byte) (*cipher.StreamReader, error){
	block, err := aes.NewCipher(key)
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])
	cipherReader := &cipher.StreamReader{S: stream, R: reader}
	return cipherReader, err	
}

func GetGZipAESWriter(writer io.Writer, key []byte) (*gzip.Writer, error) {
	if cipherWriter, err := GetAESEncryptionWriter(writer, key); err != nil {
		return nil, err
	} else {
		return gzip.NewWriter(cipherWriter), nil
	}
}

func GetGZipAESReader(reader io.Reader, key []byte) (*gzip.Reader, error) {
	if cipherReader, err := GetAESEncryptionReader(reader, key); err != nil {
		return nil, err
	} else {
		return gzip.NewReader(cipherReader)
	}
}

func GetFileMetadataHash(file string) (uint32, error) {
	if stat, err := os.Stat(file); err != nil {
		return 0, err
	} else {
		return hashutil.GetHashInt(stat), nil
	}
}