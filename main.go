package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/MC-Dashify/launcher/config"
	"github.com/MC-Dashify/launcher/global"
	"github.com/MC-Dashify/launcher/i18n"
	"github.com/MC-Dashify/launcher/rest"
	"github.com/MC-Dashify/launcher/traffic"
	"github.com/MC-Dashify/launcher/utils"
	"github.com/MC-Dashify/launcher/utils/logger"
	"github.com/MC-Dashify/launcher/webconsole"
	"github.com/gin-gonic/gin"

	"github.com/cavaliergopher/grab/v3"
)

var serverFilePath string

type downloadResult struct {
	file  string
	dlerr error
}

type runtimeJar struct {
	serverFile string
	arguments  []string
}

func parseFlags() {
	langFlag := flag.String("lang", "", i18n.Get("flag.lang.desc"))
	verboseFlag := flag.Bool("verbose", false, i18n.Get("flag.verbose.desc"))
	mcoriginFlag := flag.Int("mcorigin", 25565, i18n.Get("flag.mcorigin.desc"))
	versionFlag := flag.Bool("version", false, i18n.Get("flag.version.desc"))
	configHelpFlag := flag.Bool("config-help", false, i18n.Get("flag.config.help.desc"))

	flag.Parse()
	if (*langFlag) != "" {
		global.IsLanguageForced = true
		global.ForcedLanguage = *langFlag
	}
	if *versionFlag {
		logger.Info(strings.ReplaceAll(i18n.Get("version.info"), "$version", global.Version))
		os.Exit(0)
	}
	if *configHelpFlag {
		logger.Info(i18n.Get("config.help"))
		os.Exit(0)
	}
	if *verboseFlag {
		global.IsVerbose = true
	}
	if (*mcoriginFlag) != 25565 {
		global.MCOriginPort = *mcoriginFlag
	}

}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	logger.InitLogger()
	parseFlags()
}

func main() {
	config.ConfigContent = config.LoadConfig()

	if config.ConfigContent.EnableTrafficMonitor {
		logger.Info(strings.ReplaceAll(i18n.Get("traffic.monitor.enabled"), "$redirectPort", fmt.Sprint(config.ConfigContent.TrafficRedirectPort)))
		go traffic.StartTrafficMonitor()
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		errorMessage := "No Error."
		if param.ErrorMessage != "" {
			errorMessage = param.ErrorMessage
		}
		logger.Debug(fmt.Sprintf("[REST] [%s] | Request Method: [%s] | Request Path: [%s] | Request Protocol: [%s] | Request Status Code: [%d] | User-Agent: [%s] | Error Message: \"%s\"\n",
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Request.UserAgent(),
			errorMessage,
		))
		return ""
	}))
	router.Use(gin.Recovery())
	router.Use(rest.Authorization())
	router.Use(rest.Cors())

	router.GET("/console", webconsole.HandleWebSocket)
	router.GET("/ping", rest.Ping)
	router.GET("/logs", rest.Logs)
	router.GET("/traffic", rest.Traffic)

	router.GET("/", rest.ReverseProxy())
	router.GET("/worlds", rest.ReverseProxy())
	router.GET("/worlds/:uuid", rest.ReverseProxy())
	router.GET("/players", rest.ReverseProxy())
	router.GET("/players/:uuid", rest.ReverseProxy())
	router.POST("/players/:uuid/kick", rest.ReverseProxy())
	router.POST("/players/:uuid/ban", rest.ReverseProxy())
	router.GET("/stats", rest.ReverseProxy())

	webconsole.Server = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.ConfigContent.APIPort),
		Handler: router, // gin 핸들러 사용
	}
	go webconsole.Server.ListenAndServe()

	config.GetPluginConfig()

	runner()
	for {
		if !global.NormalStatusExit {
			logger.Fatal(i18n.Get("general.server.crashed"))
		}

		if !global.IsMCServerRunning {
			if config.ConfigContent.Restart {
				logger.Info(i18n.Get("general.server.restart"))
				fmt.Print("> ")
				time.Sleep(5 * time.Second)
				webconsole.IsRestart = true
				fmt.Print("\n")
				runner()
			} else {
				logger.Info(i18n.Get("general.exiting"))
				webconsole.Server.Shutdown(context.Background())
				os.Exit(0)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func runner() {
	javaFlavour, javaVersion := utils.CheckJava()
	logger.Info(strings.ReplaceAll(strings.ReplaceAll(i18n.Get("java.detected"), "$javaFlavour", javaFlavour), "$javaVersion", javaVersion))

	if strings.HasPrefix(config.ConfigContent.Server, "http://") || strings.HasPrefix(config.ConfigContent.Server, "https://") {
		dlServerChannel := make(chan bool)

		go downloadJar([]string{config.ConfigContent.Server}, "server", dlServerChannel)

		dlServerResult := <-dlServerChannel

		if dlServerResult {
			logger.Info(i18n.Get("general.server.download.success"))
		} else {
			logger.Error(i18n.Get("general.server.download.failed"))
		}
	} else if strings.HasPrefix(config.ConfigContent.Server, "file") {
		logger.Info(i18n.Get("general.server.source.local"))
		serverFilePath = strings.ReplaceAll(config.ConfigContent.Server, "file://", "")
		if _, err := os.Stat(serverFilePath); os.IsNotExist(err) {
			logger.Fatal(i18n.Get("general.server.source.local.notfound.or.permission.denied"))
		}
	} else {
		logger.Fatal(i18n.Get("general.server.source.invalid.protocol"))
	}

	if len(config.ConfigContent.Plugins) > 0 {
		dlPluginsChannel := make(chan bool)
		go downloadJar(config.ConfigContent.Plugins, "plugins", dlPluginsChannel)
		dlPluginsResult := <-dlPluginsChannel

		if dlPluginsResult {
			logger.Info(i18n.Get("general.plugin.download.success"))
		} else {
			logger.Error(i18n.Get("general.plugin.download.failed"))
		}
	} else {
		logger.Info(i18n.Get("general.plugin.empty"))
	}

	runtimeArgs := prepareRuntime(runtimeJar{}, config.ConfigContent)

	customArgs := append(append(runtimeArgs.arguments, "-jar"), runtimeArgs.serverFile)

	for _, customArg := range config.ConfigContent.JarArgs {
		customArgs = append(customArgs, customArg)
	}
	global.JarArgs = customArgs
	startServer(customArgs)

	// <-done
}

func startServer(customArgs []string) {
	logger.Info(i18n.Get("general.server.starting"))
	webconsole.RunServer(customArgs)
}

func downloadJar(urls []string, downloadType string, complete chan<- bool) {
	downloadDest := ""
	results := make(map[string]error)
	dlChannel := make(chan downloadResult)

	if downloadType == "server" {
		currentDir, err := os.Getwd()
		if err != nil {
			logger.Fatal(strings.ReplaceAll(i18n.Get("general.cwd.get.failed"), "$error", err.Error()))
		}
		serverDirectory := ""
		if runtime.GOOS == "windows" {
			serverDirectory = currentDir + "\\.launcher\\jars\\"
		} else {
			serverDirectory = currentDir + "/.launcher/jars/"
		}

		logger.Info(strings.ReplaceAll(i18n.Get("general.checking.directory"), "$dir", downloadType))
		utils.CheckIsExist(serverDirectory)
		downloadDest = serverDirectory

	} else if downloadType == "plugins" {
		currentPath, _ := os.Getwd()
		pluginDirectory := currentPath + "/plugins/"

		logger.Info(strings.ReplaceAll(i18n.Get("general.checking.directory"), "$dir", downloadType))
		utils.CheckIsExist(pluginDirectory)
		downloadDest = pluginDirectory

	} else {
		logger.Fatal(i18n.Get("general.download.type.invalid"))
	}

	logger.Info(strings.ReplaceAll(i18n.Get("general.download.preparing"), "$file", downloadType))
	for _, url := range urls {
		go downloadFile(downloadType, downloadDest, url, dlChannel)
	}

	for i := 0; i < len(urls); i++ {
		downloadResult := <-dlChannel
		results[downloadResult.file] = downloadResult.dlerr
	}

	for downloadedFile, downloadError := range results {
		if downloadError != nil {
			logger.Error(strings.ReplaceAll(strings.ReplaceAll(i18n.Get("general.download.failed"), "$file", downloadedFile), "$error", downloadError.Error()))
		}
	}
	logger.Info(strings.ReplaceAll(i18n.Get("general.download.success"), "$type", downloadType))
	complete <- true
}

func downloadFile(downloadType, downloadDir, url string, err chan<- downloadResult) {
	if downloadType == "server" && !utils.IsValidUrl(url) {
		err <- downloadResult{file: url, dlerr: nil}
		return
	}

	client := grab.NewClient()
	req, _ := grab.NewRequest(downloadDir, url)
	req.NoResume = false
	resp := client.Do(req)

	t := time.NewTicker(time.Second)
	defer t.Stop()
Loop:
	for {
		select {
		case <-t.C:
			etaTime := time.Until(resp.ETA()).Round(time.Second).String()

			if strings.Contains(etaTime, "-") {
				etaTime = i18n.Get("general.calculating")
			}

			downloadSpeed := utils.ByteCounter(int64(resp.BytesPerSecond()))
			currentDownloaded := utils.ByteCounter(resp.BytesComplete())
			totalDownloaded := utils.ByteCounter(resp.Size())

			var jarPath []string
			if runtime.GOOS == "windows" {
				jarPath = strings.Split(resp.Filename, "\\")
			} else {
				jarPath = strings.Split(resp.Filename, "/")
			}
			logger.Info(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(i18n.Get("general.download.progress"), "$fileName", jarPath[len(jarPath)-1]), "$downloadedSize", currentDownloaded), "$fileSize", totalDownloaded), "$eta", etaTime), "$downloadSpeed", downloadSpeed))

		case <-resp.Done:
			break Loop
		}
	}

	if dlErr := resp.Err(); dlErr != nil {
		logger.Error(strings.ReplaceAll(strings.ReplaceAll(i18n.Get("general.download.failed"), "$file", resp.Filename), "$error", dlErr.Error()))
		err <- downloadResult{file: resp.Filename, dlerr: dlErr}
	}

	jarPath := strings.Split(resp.Filename, "/")

	logger.Info(strings.ReplaceAll(i18n.Get("general.download.done"), "$fileName", jarPath[len(jarPath)-1]))

	if downloadType == "server" {
		serverFilePath = resp.Filename
	}
	err <- downloadResult{file: resp.Filename, dlerr: nil}
}

func prepareRuntime(runtime runtimeJar, config config.Config) runtimeJar {
	runtime = runtimeJar{serverFile: serverFilePath}

	for _, option := range []string{
		fmt.Sprintf("-Xmx%dG", config.Memory),
		fmt.Sprintf("-Xms%dG", config.Memory),
		"-XX:+ParallelRefProcEnabled",
		"-XX:MaxGCPauseMillis=200",
		"-XX:+UnlockExperimentalVMOptions",
		"-XX:+DisableExplicitGC",
		"-XX:+AlwaysPreTouch",
		"-XX:G1HeapWastePercent=5",
		"-XX:G1MixedGCCountTarget=4",
		"-XX:G1MixedGCLiveThresholdPercent=90",
		"-XX:G1RSetUpdatingPauseTimePercent=5",
		"-XX:SurvivorRatio=32",
		"-XX:+PerfDisableSharedMem",
		"-XX:MaxTenuringThreshold=1",
		"-Dusing.aikars.flags=https://mcflags.emc.gs",
		"-Daikars.new.flags=true",
		"-Dcom.mojang.eula.agree=true",
	} {
		runtime.arguments = append(runtime.arguments, option)
	}
	for _, option := range utils.SelectOptionByMemory(config.Memory) {
		runtime.arguments = append(runtime.arguments, option)
	}

	if config.Debug {
		debugOption := "-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address="
		_, javaVersion := utils.CheckJava()
		if utils.VersionOrdinal("1.8") < utils.VersionOrdinal(javaVersion) {
			debugOption += fmt.Sprintf("*:%d", config.DebugPort)
		} else {
			debugOption += fmt.Sprintf("%d", config.DebugPort)
		}
		runtime.arguments = append(runtime.arguments, debugOption)
	}

	return runtime
}
