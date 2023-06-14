package rest

import (
	"fmt"
	"net/http"

	"github.com/MC-Dashify/launcher/config"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/console" {
			websocketAuth := c.Request.URL.Query().Get("auth_key")
			if websocketAuth == config.GetPluginConfig().Key {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
				c.Abort()
			}
			return
		}
		authKey := c.Request.Header.Get("Authorization")
		if authKey == fmt.Sprintf("Bearer %s", config.GetPluginConfig().Key) {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			c.Abort()
		}
	}
}
