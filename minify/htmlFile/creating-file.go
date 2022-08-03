package htmlFile

import (
	"errors"
	"os"
	"strings"
)

func MakeFilePath() (string, error) {
	filePath, err := GetFilePath()
	if err != nil {
		return "", err
	}
	filePath = strings.Replace(filePath, ".html", "_Copy.html", 1)
	return filePath, nil
}

func Create() (*os.File, error) {
	filePath, err := MakeFilePath()
	if err != nil {
		return nil, err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return nil, errors.New("Can't Create File " + err.Error())
	}
	return file, nil
}
