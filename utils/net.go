package utils

import (
	"net/http"
	"net/url"
	"time"

	"github.com/MC-Dashify/launcher/utils/logger"
)

func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func GetLastModifiedFromUrl(url string) int64 {
	resp, err := http.Head(url)

	if err != nil {
		logger.Error("Failed to fetch file info from url")
	} else {
		defer resp.Body.Close()
		_remoteFileTime, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))
		if err != nil {
			logger.Warn("Failed to parse time")
		} else {
			return _remoteFileTime.Unix()
		}
	}
	return time.Now().Unix()
}
