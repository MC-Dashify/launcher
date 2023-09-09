package rest

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type DirectoryStatus struct {
	FileName    string
	IsDirectory bool
}

func GETFiles(c *gin.Context) {
	var err error
	requestURL := c.Request.URL.Path
	fsPath := strings.Replace(requestURL, "/files", ".", 1)
	_executablePath, err := os.Executable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": fsPath})
		return
	}
	_absFSPath := filepath.Join(filepath.Dir(_executablePath), fsPath)

	// first, check if the requested path is in the root directory
	_rootDirPath := filepath.Dir(_executablePath)
	_targetFilePath := _absFSPath

	if !strings.HasPrefix(_targetFilePath, _rootDirPath) {
		_targetFilePath = _rootDirPath
	}

	if filepath.ToSlash(_executablePath) == filepath.ToSlash(_targetFilePath) {
		c.JSON(http.StatusNotAcceptable, gin.H{"status": "failed", "detail": "cannot modify launcher itself.", "path": fsPath})
		return
	}

	// then, check if the requested path is exists
	if _, err := os.Stat(_targetFilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "detail": err.Error(), "path": fsPath})
		return
	}

	// then check if the requested path is a directory
	fileInfo, err := os.Stat(_targetFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": fsPath})
		return
	}

	if fileInfo.IsDir() {
		files, err := os.ReadDir(_targetFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": fsPath})
			return
		}
		dirStatus := []DirectoryStatus{}
		for _, file := range files {
			if (file.Name() != filepath.Base(_executablePath)) && (filepath.ToSlash(_executablePath) != filepath.ToSlash(_targetFilePath)) {
				dirStatus = append(dirStatus, DirectoryStatus{FileName: file.Name(), IsDirectory: file.IsDir()})
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "content": dirStatus, "path": fsPath})
	} else {
		c.Status(http.StatusOK)
		c.File(_targetFilePath)
	}
}
