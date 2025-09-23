package utils

import (
	"os"
	"path/filepath"
)

func CreateFolder(name string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(cwd, name)
	return os.MkdirAll(path, 0o755) // 755 safe permission
}

// create an empty file, overwrites if already exist, use with cautions
func CreateFile(name string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(cwd, name)
	_, err = os.Create(path)
	return err
}

func CreateFileWithContent(name, content string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(cwd, name)
	return os.WriteFile(path, []byte(content), 0o644) // 644 for files
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
