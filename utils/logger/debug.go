package logger

import (
	"fmt"
	"time"

	"github.com/MC-Dashify/launcher/global"
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
	if global.IsVerbose {
		customLog.SetLevel(customLog.DebugLevel)
	}
	customLog.Debugf(content)
}
