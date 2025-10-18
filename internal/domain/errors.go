package domain

import "errors"

var (
	ErrUnsupportedDB   = errors.New("unsupported database driver")
	ErrUnsupportedLang = errors.New("unsupported language driver")
	ErrInvalidSchema   = errors.New("invalid schema file")
)
