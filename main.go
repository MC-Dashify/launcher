package main

import (
	"flag"
	"strings"

	"github.com/MC-Dashify/launcher/global"
	"github.com/MC-Dashify/launcher/i18n"
	"github.com/MC-Dashify/launcher/utils/logger"
)

func parseFlags() {
	wordPtr := flag.String("lang", "", i18n.Get("flag.lang.desc"))

	flag.Parse()
	if (*wordPtr) != "" {
		global.IsLanguageForced = true
		global.ForcedLanguage = *wordPtr
	}
}

func init() {
	logger.InitLogger()
	parseFlags()
}

func main() {
	logger.Info(i18n.Get("test.hello"))
	logger.Info(strings.ReplaceAll(i18n.Get("test.placeholder"), "$user", "testuser1"))
}
