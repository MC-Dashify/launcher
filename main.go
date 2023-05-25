package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/MC-Dashify/launcher/global"
	"github.com/MC-Dashify/launcher/i18n"
	"github.com/MC-Dashify/launcher/utils"
	"github.com/MC-Dashify/launcher/utils/logger"
)

func parseFlags() {
	langFlag := flag.String("lang", "", i18n.Get("flag.lang.desc"))
	verboseFlag := flag.Bool("verbose", false, i18n.Get("flag.verbose.desc"))

	flag.Parse()
	if (*langFlag) != "" {
		global.IsLanguageForced = true
		global.ForcedLanguage = *langFlag
	}
	if *verboseFlag {
		global.IsVerbose = true
	}
}

func init() {
	logger.InitLogger()
	parseFlags()
}

func main() {
	javaFlavour, javaVersion := utils.CheckJava()
	logger.Info(strings.ReplaceAll(strings.ReplaceAll(i18n.Get("java.detected"), "$javaFlavour", javaFlavour), "$javaVersion", javaVersion))
	logger.Info(fmt.Sprintf("%d", utils.GetLastModified("https://clip.aroxu.me/download?mc_version=1.19.4")))
}
