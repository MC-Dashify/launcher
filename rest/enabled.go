package rest

import (
	"github.com/MC-Dashify/launcher/config"
	"github.com/gin-gonic/gin"
)

func IsEnabled() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !config.GetPluginConfig().Enabled {
			c.AbortWithStatusJSON(418, gin.H{"status": "I'm a tea pot :3", "detail": "server disabled plugin."})
		}
	}
}
