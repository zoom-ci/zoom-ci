// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gopath

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CurrentPath() (string, error) {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

func CurrentParentPath() (string, error) {
	path := strings.Join([]string{filepath.Dir(os.Args[0]), "/../"}, "")
	realPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return realPath, nil
}

func CreatePath(path string) error {
	exists := Exists(path)
	if !exists {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsExist(err) {
		return true
	}
	return false
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func RemovePath(path string) error {
	return os.RemoveAll(path)
}

func SaveFile(fileName string, fileContent []byte) error {
	fileExists := Exists(fileName)
	if !fileExists {
		return ioutil.WriteFile(fileName, fileContent, fs.FileMode.Perm(0777))
	}
	return nil
}