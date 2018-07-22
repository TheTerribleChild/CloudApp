package fileutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

//GetAllFileInDirectoryRecursively Gets all regular file with extension in root directory recursively
func GetAllFileInDirectoryRecursively(root string, ext string) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f.Mode().IsRegular() {
			if (len(ext) != 0 && filepath.Ext(path) == ext) || len(ext) == 0 {
				fileList = append(fileList, path)
			}
		}
		return nil
	})
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
