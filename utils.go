package viperjacket

import (
	"path/filepath"
)

func getFileDir(path string) string {
	return filepath.Dir(path)
}

func getFileName(path string) string {
	return filepath.Base(path)
}

func getFileExtension(path string) string {
	ext := filepath.Ext(path)
	if len(ext) > 0 {
		return ext[1:]
	}
	return ext
}
