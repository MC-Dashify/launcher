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
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
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
		c.JSON(http.StatusNotAcceptable, gin.H{"status": "failed", "detail": "cannot read launcher itself.", "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	// then, check if the requested path is exists
	if _, err := os.Stat(_targetFilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	// then check if the requested path is a directory
	fileInfo, err := os.Stat(_targetFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	if fileInfo.IsDir() {
		files, err := os.ReadDir(_targetFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
			return
		}
		dirStatus := []DirectoryStatus{}
		for _, file := range files {
			if (file.Name() != filepath.Base(_executablePath)) && (filepath.ToSlash(_executablePath) != filepath.ToSlash(_targetFilePath)) {
				dirStatus = append(dirStatus, DirectoryStatus{FileName: file.Name(), IsDirectory: file.IsDir()})
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "content": dirStatus, "path": strings.Replace(fsPath, ".", "", 1)})
	} else {
		c.Status(http.StatusOK)
		c.File(_targetFilePath)
	}
}

func POSTFiles(c *gin.Context) {
	var err error
	requestURL := c.Request.URL.Path
	fsPath := strings.Replace(requestURL, "/files", ".", 1)
	_executablePath, err := os.Executable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}
	_absFSPath := filepath.Join(filepath.Dir(_executablePath), fsPath)

	// first, check if the requested path is in the root directory
	_rootDirPath := filepath.Dir(_executablePath)
	_targetFilePath := _absFSPath

	if !strings.HasPrefix(_targetFilePath, _rootDirPath) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "detail": "cannot create file outside of server.", "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	if filepath.ToSlash(_executablePath) == filepath.ToSlash(_targetFilePath) {
		c.JSON(http.StatusNotAcceptable, gin.H{"status": "failed", "detail": "cannot read launcher itself.", "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	// then, check if the requested path is exists
	if _, err := os.Stat(_targetFilePath); !os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "detail": "Requested path is already exist.", "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		if strings.HasSuffix(fsPath, "/") {
			os.MkdirAll(_targetFilePath, os.ModePerm)
			c.JSON(http.StatusOK, gin.H{"status": "ok", "path": strings.Replace(fsPath, ".", "", 1)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	if file.Size < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "detail": "No files uploaded.", "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	err = c.SaveUploadedFile(file, _targetFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "path": strings.Replace(fsPath, ".", "", 1)})
}

func PATCHFiles(c *gin.Context) {
	var err error
	requestURL := c.Request.URL.Path
	fsPath := strings.Replace(requestURL, "/files", ".", 1)
	_executablePath, err := os.Executable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}
	_absFSPath := filepath.Join(filepath.Dir(_executablePath), fsPath)

	// first, check if the requested path is in the root directory
	_rootDirPath := filepath.Dir(_executablePath)
	_targetFilePath := _absFSPath

	if !strings.HasPrefix(_targetFilePath, _rootDirPath) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "detail": "cannot modify file outside of server.", "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	if filepath.ToSlash(_executablePath) == filepath.ToSlash(_targetFilePath) {
		c.JSON(http.StatusNotAcceptable, gin.H{"status": "failed", "detail": "cannot read launcher itself.", "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	// then, check if the requested path is exists
	if _, err := os.Stat(_targetFilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	if file.Size < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "detail": "No files uploaded.", "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	err = c.SaveUploadedFile(file, _targetFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "path": strings.Replace(fsPath, ".", "", 1)})
}

type MoveFile struct {
	MoveTo string
}

func PUTFiles(c *gin.Context) {
	var err error
	requestURL := c.Request.URL.Path
	originalFilePath := strings.Replace(requestURL, "/files", ".", 1)
	var moveFilePathObject MoveFile

	if err := c.BindJSON(&moveFilePathObject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "detail": err.Error(), "path": originalFilePath})
	}

	_executablePath, err := os.Executable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": originalFilePath})
		return
	}
	_absOriginalFSPath := filepath.Join(filepath.Dir(_executablePath), originalFilePath)
	_absMoveFSPath := filepath.Join(filepath.Dir(_executablePath), moveFilePathObject.MoveTo)

	// first, check if the requested path is in the root directory
	_rootDirPath := filepath.Dir(_executablePath)
	_targetFilePath := _absOriginalFSPath
	_targetMoveFilePath := _absMoveFSPath

	if !strings.HasPrefix(_targetFilePath, _rootDirPath) || !strings.HasPrefix(_targetMoveFilePath, _rootDirPath) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "detail": "cannot modify file outside of server.", "path": originalFilePath})
		return
	}

	if filepath.ToSlash(_executablePath) == filepath.ToSlash(_targetFilePath) || filepath.ToSlash(_executablePath) == filepath.ToSlash(_targetMoveFilePath) {
		c.JSON(http.StatusNotAcceptable, gin.H{"status": "failed", "detail": "cannot read launcher itself.", "path": originalFilePath})
		return
	}

	// then, check if the requested path is exists
	if _, err := os.Stat(_targetFilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "detail": err.Error(), "path": originalFilePath})
		return
	}
	if _, err := os.Stat(_targetMoveFilePath); os.IsExist(err) {
		c.JSON(http.StatusConflict, gin.H{"status": "failed", "detail": err.Error(), "path": moveFilePathObject.MoveTo})
		return
	}

	err = os.Rename(_absOriginalFSPath, _absMoveFSPath)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "origin_path": strings.Replace(originalFilePath, ".", "", 1), "moved_path": moveFilePathObject.MoveTo})
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "origin_path": originalFilePath, "moved_path": moveFilePathObject.MoveTo})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"status": "ok", "origin_path": strings.Replace(originalFilePath, ".", "", 1), "moved_path": moveFilePathObject.MoveTo})
	c.JSON(http.StatusOK, gin.H{"status": "ok", "origin_path": originalFilePath, "moved_path": moveFilePathObject.MoveTo})
}

func DELETEFiles(c *gin.Context) {
	var err error
	requestURL := c.Request.URL.Path
	fsPath := strings.Replace(requestURL, "/files", ".", 1)
	_executablePath, err := os.Executable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
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
		c.JSON(http.StatusNotAcceptable, gin.H{"status": "failed", "detail": "cannot delete launcher itself.", "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	// then, check if the requested path is exists
	if _, err := os.Stat(_targetFilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	// then delete requested path
	err = os.RemoveAll(_targetFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "detail": err.Error(), "path": strings.Replace(fsPath, ".", "", 1)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "path": strings.Replace(fsPath, ".", "", 1)})
}
