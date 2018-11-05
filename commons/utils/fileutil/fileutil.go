package fileutil

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"path"
)

//GetAllFileInDirectoryRecursively Gets all regular file with extension in root directory recursively
func GetAllFileInDirectoryRecursively(roots []string, ext string) ([]string, error) {
	fileList := []string{}
	var err error
	for _, root := range roots {
		err = filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
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

func ZipFiles(files []string, outputFilePath string) error {
	if len(outputFilePath) == 0 {
		return fmt.Errorf("No output file path given.")
	}
	
	commonPrefix := GetCommonPrefix(files)
	removeIdx := len(commonPrefix)

	newZipFile, err := os.Create(outputFilePath)
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
			if _, err := io.CopyN(writer, zipfile, 1048576); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
		}
		zipfile.Close()
	}
	return nil
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