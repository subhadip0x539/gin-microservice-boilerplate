package utils

import (
	"fmt"
	"go-gin-boilerplate/api/logging"
	"os"
	"path/filepath"
)

func CreateDirectories(directoryNames []string) {
	baseDirectory, error := os.Getwd()
	if error != nil {
		fmt.Println(error)
	}

	for _, directoryName := range directoryNames {
		directoryPath := filepath.Join(baseDirectory, directoryName)
		if _, error := os.Stat(directoryPath); os.IsNotExist(error) {
			error := os.MkdirAll(directoryPath, 0755)
			if error != nil {
				logging.Logger.Error(error.Error())
			}
		}
	}

}
