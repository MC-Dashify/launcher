package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/MC-Dashify/launcher/i18n"
	"github.com/MC-Dashify/launcher/utils/logger"
)

var ConfigContent Config

const configTemplate = `{
  "server": "https://clip.aroxu.me/download?mc_version=1.19.4",
  "debug": false,
  "debug_port": 5005,
  "restart": true,
  "memory": 4,
	"api_port": 8080,
	"plugin_api_port": 8081,
  "plugins": [
    "https://github.com/monun/auto-reloader/releases/download/0.0.6/auto-reloader-0.0.6.jar"
  ],
  "jarArgs": ["nogui"]
}`

type Config struct {
	Server     string   `json:"server"`
	Debug      bool     `json:"debug"`
	DebugPort  int      `json:"debug_port"`
	Restart    bool     `json:"restart"`
	Memory     int      `json:"memory"`
	APIPort    int      `json:"api_port"`
	PluginPort int      `json:"plugin_api_port"`
	Plugins    []string `json:"plugins"`
	JarArgs    []string `json:"jarArgs"`
}

func LoadConfig() Config {
	var config Config
	currentPath, _ := os.Getwd()
	configPath := currentPath + "/launcher.conf.json"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logger.Warn(i18n.Get("config.notfound"))
		generateConfig()
	}

	configData, loadFileErr := os.ReadFile(configPath)
	if loadFileErr != nil {
		logger.Error(strings.ReplaceAll(i18n.Get("config.loaderror"), "$error", loadFileErr.Error()))
	}
	if strings.TrimSpace(string(configData)) == "" {
		logger.Error(i18n.Get("config.empty"))
		configData = []byte(configTemplate)
		generateConfig()
	}

	loadConfigErr := json.Unmarshal([]byte(configData), &config)
	if loadConfigErr != nil {
		logger.Fatal(strings.ReplaceAll(i18n.Get("config.invalid"), "$error", loadConfigErr.Error()))
	} else {
		if config.Server == "" {
			logger.Warn(i18n.Get("config.server.empty"))
		}
		if config.APIPort <= 0 || config.APIPort > 65535 {
			logger.Fatal(i18n.Get("config.api_port.invalid"))
		}
		if config.PluginPort <= 0 || config.PluginPort > 65535 {
			logger.Fatal(i18n.Get("config.plugin_api_port.invalid"))
		}
	}
	return config
}

func generateConfig() {
	serverConfFile, errGenConf := os.Create("launcher.conf.json")

	if errGenConf != nil {
		logger.Fatal(strings.ReplaceAll(i18n.Get("config.create_failed"), "$error", errGenConf.Error()))
	}

	defer serverConfFile.Close()

	_, errWrtConf := serverConfFile.WriteString(configTemplate)

	if errWrtConf != nil {
		logger.Fatal(strings.ReplaceAll(i18n.Get("config.write_failed"), "$error", errWrtConf.Error()))
	}
	logger.Info(i18n.Get("config.created"))
}
