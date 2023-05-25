package logger

import (
	"fmt"
	"time"

	customLog "github.com/sirupsen/logrus"
)

// Error logs error level
func Error(content string) {
	content = fmt.Sprintf("%v ", time.Now().Format("2006-01-02 15:04:05")) + content

	customLog.SetFormatter(&customLog.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
	})
	customLog.SetOutput(ColorableStdout)
	customLog.SetLevel(customLog.ErrorLevel)
	customLog.Errorf(content)
}
