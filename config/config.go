package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/MC-Dashify/launcher/i18n"
	"github.com/MC-Dashify/launcher/utils/logger"
)

var ConfigContent Config

const configFileName = "launcher.conf.json"

const (
	defaultConfigVersion = 1
	defaultServer        = "https://clip.aroxu.me/download?mc_version=1.19.4"
	defaultDebug         = false
	defaultDebugPort     = 5005
	defaultRestart       = true
	defaultMemory        = 2
	defaultAPIPort       = 8080
	defaultPluginPort    = 8081
)

var (
	defaultPlugins                = []string{}
	defaultJarArgs                = []string{"nogui"}
	defaultWebConsoleDisabledCmds = []string{}
)

type Config struct {
	ConfigVersion          int      `json:"config_version"`
	Server                 string   `json:"server"`
	Debug                  bool     `json:"debug"`
	DebugPort              int      `json:"debug_port"`
	Restart                bool     `json:"restart"`
	Memory                 int      `json:"memory"`
	APIPort                int      `json:"api_port"`
	PluginPort             int      `json:"plugin_api_port"`
	Plugins                []string `json:"plugins"`
	JarArgs                []string `json:"jar_args"`
	WebConsoleDisabledCmds []string `json:"webconsole_disabled_cmds"`
}

func LoadConfig() Config {
	var config Config
	defaultConfig := Config{
		ConfigVersion:          defaultConfigVersion,
		Server:                 defaultServer,
		Debug:                  defaultDebug,
		DebugPort:              defaultDebugPort,
		Restart:                defaultRestart,
		Memory:                 defaultMemory,
		APIPort:                defaultAPIPort,
		PluginPort:             defaultPluginPort,
		Plugins:                defaultPlugins,
		JarArgs:                defaultJarArgs,
		WebConsoleDisabledCmds: defaultWebConsoleDisabledCmds,
	}
	currentPath, _ := os.Getwd()
	configPath := currentPath + "/" + configFileName

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logger.Warn(i18n.Get("config.notfound"))
		saveConfig(defaultConfig)
	}

	configData, loadFileErr := os.ReadFile(configPath)
	if loadFileErr != nil {
		logger.Error(strings.ReplaceAll(i18n.Get("config.loaderror"), "$error", loadFileErr.Error()))
	}

	if strings.TrimSpace(string(configData)) == "" {
		logger.Error(i18n.Get("config.empty"))
		saveConfig(defaultConfig)
		return defaultConfig
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
		if config.DebugPort <= 0 || config.DebugPort > 65535 {
			logger.Fatal(i18n.Get("config.debug_port.invalid"))
		}
		if config.Memory < 2 {
			logger.Fatal(i18n.Get("config.memory.invalid"))
		}
		if config.ConfigVersion != defaultConfigVersion {
			logger.Warn(i18n.Get("config.version.different"))
		}
	}
	return config
}

func saveConfig(data Config) {
	file, _ := json.MarshalIndent(data, "", " ")

	serverConfFile, errGenConf := os.Create(configFileName)

	if errGenConf != nil {
		logger.Fatal(strings.ReplaceAll(i18n.Get("config.create_failed"), "$error", errGenConf.Error()))
	}

	defer serverConfFile.Close()

	_, errWrtConf := serverConfFile.Write(file)

	if errWrtConf != nil {
		logger.Fatal(strings.ReplaceAll(i18n.Get("config.write_failed"), "$error", errWrtConf.Error()))
	}
	logger.Info(i18n.Get("config.created"))
}
