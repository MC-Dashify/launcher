package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/MC-Dashify/launcher/utils/logger"
)

func CheckFolderExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Warn("Generating missing folders")
		Mkdir(path, 0755)
	}
}

func Mkdir(path string, perm os.FileMode) {
	err := os.MkdirAll(path, perm)
	if err != nil {
		logger.Panic(err.Error())
	}
}

func ByteCounter(b int64) string {
	if b == -1 {
		return "Unknown Size"
	}
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func GetLastModifiedFromLocal(path string) int64 {
	fileinfo, err := os.Stat(path)

	if err != nil {
		logger.Error("Failed to fetch file info from local")
	} else {
		return fileinfo.ModTime().Unix()
	}
	return time.Now().Unix()
}
