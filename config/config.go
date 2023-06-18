package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/MC-Dashify/launcher/i18n"
	"github.com/MC-Dashify/launcher/utils"
	"github.com/MC-Dashify/launcher/utils/logger"
	"gopkg.in/yaml.v3"
)

var ConfigContent Config

const configFileName = "launcher.conf.json"
const pluginConfigFileName = "config.yml"

const (
	defaultConfigVersion        = 1
	defaultServer               = "https://clip.aroxu.me/download?mc_version=1.19.4"
	defaultDebug                = false
	defaultDebugPort            = 5005
	defaultEnableTrafficMonitor = true
	defaultRestart              = true
	defaultMemory               = 2
	defaultTrafficRedirectPort  = 25565
	defaultAPIPort              = 8080
	defaultPluginPort           = 8081
)

var (
	defaultPlugins                = []string{"https://github.com/MC-Dashify/plugin/releases/latest/download/dashify-plugin-all.jar"}
	defaultJarArgs                = []string{"nogui"}
	defaultWebConsoleDisabledCmds = []string{"stop"}
)

type Config struct {
	ConfigVersion          int      `json:"config_version"`
	Server                 string   `json:"server"`
	Debug                  bool     `json:"debug"`
	DebugPort              int      `json:"debug_port"`
	Restart                bool     `json:"restart"`
	Memory                 int      `json:"memory"`
	EnableTrafficMonitor   bool     `json:"enable_traffic_monitor"`
	TrafficRedirectPort    int      `json:"traffic_redirect_port"`
	APIPort                int      `json:"api_port"`
	PluginPort             int      `json:"plugin_api_port"`
	Plugins                []string `json:"plugins"`
	JarArgs                []string `json:"jar_args"`
	WebConsoleDisabledCmds []string `json:"webconsole_disabled_cmds"`
}

type PluginConfig struct {
	Key string `yaml:"key"`
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
		TrafficRedirectPort:    defaultTrafficRedirectPort,
		APIPort:                defaultAPIPort,
		PluginPort:             defaultPluginPort,
		Plugins:                defaultPlugins,
		JarArgs:                defaultJarArgs,
		WebConsoleDisabledCmds: defaultWebConsoleDisabledCmds,
	}
	currentPath, _ := os.Getwd()
	configPath := filepath.Join(currentPath, configFileName)

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
		if config.TrafficRedirectPort <= 0 || config.TrafficRedirectPort > 65535 {
			logger.Fatal(i18n.Get("config.api_port.invalid"))
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

func GetPluginConfig() PluginConfig {
	var pluginConfig PluginConfig
	currentPath, _ := os.Getwd()
	configPath := filepath.Join(currentPath, "plugins", "Dashify", pluginConfigFileName)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		randomString := utils.GenerateRandomString(64)
		hashedString := utils.GenerateBCryptString(randomString)
		pluginConfig = PluginConfig{
			Key: hashedString,
		}
		savePluginConfig(pluginConfig)
	}

	configData, loadFileErr := os.ReadFile(configPath)
	if loadFileErr != nil {
		logger.Error(loadFileErr.Error())
	}

	if strings.TrimSpace(string(configData)) == "" {
		randomString := utils.GenerateRandomString(64)
		hashedString := utils.GenerateBCryptString(randomString)
		pluginConfig = PluginConfig{
			Key: hashedString,
		}
		savePluginConfig(pluginConfig)
	}

	loadConfigErr := yaml.Unmarshal([]byte(configData), &pluginConfig)
	if loadConfigErr != nil {
		logger.Fatal(strings.ReplaceAll(i18n.Get("config.invalid"), "$error", loadConfigErr.Error()))
	} else {
		if pluginConfig.Key == "" {
			randomString := utils.GenerateRandomString(64)
			hashedString := utils.GenerateBCryptString(randomString)
			pluginConfig = PluginConfig{
				Key: hashedString,
			}
			savePluginConfig(pluginConfig)
		}
	}
	return pluginConfig
}

func savePluginConfig(data PluginConfig) {
	currentPath, _ := os.Getwd()
	configPath := filepath.Join(currentPath, "plugins", "Dashify", pluginConfigFileName)

	utils.CheckIsExist(filepath.Join(currentPath, "plugins", "Dashify"))

	file, _ := yaml.Marshal(data)

	serverConfFile, errGenConf := os.Create(configPath)

	if errGenConf != nil {
		logger.Fatal(errGenConf.Error())
	}

	defer serverConfFile.Close()

	_, errWrtConf := serverConfFile.Write(file)

	if errWrtConf != nil {
		logger.Fatal(errWrtConf.Error())
	}
}
