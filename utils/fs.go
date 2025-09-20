package utils

import (
	"os"
	"path/filepath"
)

func CreateFolder(name string) {}

func CreateFile(name, ext string) {}

func CreateFileWithContent(name, ext string, content []byte) {

}

func FileExists(path string) bool {
	cwd, err := os.Getwd()
	if err != nil {
		return false
	}

	fullPath := filepath.Join(cwd, path)
	info, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && !info.IsDir()
}

func FolderExists(name string) bool {
	cwd, err := os.Getwd()
	if err != nil {
		return false
	}

	path := cwd + "/" + name
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}

	return info.IsDir()
}

func GetFileContent(path string) {

}
