package utils

import (
	"os"

	"github.com/MC-Dashify/launcher/utils/logger"
)

func CheckFolderExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Warn("Generating missing folders")
		Mkdir(path, 0755)
	}
}

func Mkdir(path string, perm os.FileMode) {
	err := os.Mkdir(path, perm)
	if err != nil {
		logger.Panic(err.Error())
	}
}
