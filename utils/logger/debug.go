package logger

import (
	"fmt"
	"time"

	customLog "github.com/sirupsen/logrus"
)

// Debug logs debug level
func Debug(content string) {
	content = fmt.Sprintf("%v ", time.Now().Format("2006-01-02 15:04:05")) + content

	customLog.SetFormatter(&customLog.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
	})
	customLog.SetOutput(ColorableStdout)
	customLog.SetLevel(customLog.DebugLevel)
	customLog.Debugf(content)
}
