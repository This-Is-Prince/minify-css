package htmlFile

import (
	"errors"
	"os"
	"strings"
)

func GetFilePath() (string, error) {
	args := os.Args
	if len(args) > 1 {
		fileName := args[1]
		if strings.Contains(fileName, ".html") {
			return fileName, nil
		} else {
			return "", errors.New("no HTML file path provided")
		}
	} else {
		return "", errors.New("no cmd arguments(HTML file path) provided")
	}
}

func Open() (*os.File, error) {
	filePath, err := GetFilePath()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Can't Open File " + err.Error())
	}
	return file, nil
}
