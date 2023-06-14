package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MC-Dashify/launcher/utils"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func Logs(c *gin.Context) {
	stringLines := c.Request.URL.Query().Get("lines")
	lines, err := strconv.Atoi(stringLines)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid lines query"})
		return
	}
	if lines > 1000 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Lines query too large"})
		return
	}
	if lines < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Lines query too small"})
		return
	}
	logs, readErr := utils.ReadLastNLines(lines)
	if readErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error reading logs", "detail": fmt.Sprintf("%v", readErr)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"logs": logs})
}
