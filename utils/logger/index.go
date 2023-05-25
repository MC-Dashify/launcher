package logger

import (
	"io"

	"github.com/mattn/go-colorable"
)

var ColorableStdout io.Writer

func InitLogger() {
	ColorableStdout = colorable.NewColorableStdout()
}
