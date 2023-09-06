package rest

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/MC-Dashify/launcher/config"
	"github.com/MC-Dashify/launcher/utils/logger"
	"github.com/gin-gonic/gin"
)

func ReverseProxy() gin.HandlerFunc {
	target := fmt.Sprintf("localhost:%d", config.ConfigContent.PluginPort)

	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = target
			req.Header = c.Request.Header
			// req.Header["my-header"] = []string{r.Header.Get("my-header")}
			// // Golang camelcases headers
			// delete(req.Header, "My-Header")
		}
		proxy := &httputil.ReverseProxy{
			Director: director, ErrorHandler: func(_ http.ResponseWriter, r *http.Request, err error) {
				logger.Debug(fmt.Sprintf("%+v", err))
			},
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
