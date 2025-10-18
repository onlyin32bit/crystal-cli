package fs

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
