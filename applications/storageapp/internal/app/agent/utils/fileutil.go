package util

import(
	"path/filepath"
	"os"
)

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