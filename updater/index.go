package updater

import (
	"fmt"
	"os"
	"strings"

	"github.com/MC-Dashify/launcher/global"
	"github.com/MC-Dashify/launcher/i18n"
	"github.com/MC-Dashify/launcher/utils/logger"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func ConfirmAndSelfUpdate() {
	logger.Info(fmt.Sprintf("[Updater] %+v", i18n.Get("updater.checking")))
	latest, found, err := selfupdate.DetectLatest("MC-Dashify/launcher")
	if err != nil {
		logger.Error(fmt.Sprintf("[Updater] %+v", strings.ReplaceAll(i18n.Get("updater.version.fetch.failed"), "$error", err.Error())))
		return
	}

	v := semver.MustParse(strings.TrimSpace(global.VERSION))
	if !found || latest.Version.LTE(v) {
		logger.Info(fmt.Sprintf("[Updater] %+v", i18n.Get("updater.up.to.date")))
		return
	}

	exe, err := os.Executable()
	if err != nil {
		logger.Error(fmt.Sprintf("[Updater] %+v", i18n.Get("updater.executable.path.notfound")))
		return
	}
	logger.Info(fmt.Sprintf("[Updater] %+v", i18n.Get("updater.new.version.found")))
	if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
		logger.Error(fmt.Sprintf("[Updater] %+v", strings.ReplaceAll(i18n.Get("updater.update.failed"), "$error", err.Error())))
		return
	}
	logger.Info(fmt.Sprintf("[Updater] %+v", strings.ReplaceAll(i18n.Get("updater.update.success"), "$version", latest.Version.String())))
	logger.Info(fmt.Sprintf("[Updater] %+v", i18n.Get("updater.restart.to.apply")))
	os.Exit(0)
}
