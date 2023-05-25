package main

import "github.com/MC-Dashify/launcher/utils/logger"

func init() {
	logger.InitLogger()
}

func main() {
	logger.Debug("This is debug")
	logger.Error("This is error")
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Fatal("This is fatal")
	logger.Panic("This is panic")
}
