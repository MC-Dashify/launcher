package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/MC-Dashify/launcher/config"
	"github.com/MC-Dashify/launcher/global"
	"github.com/MC-Dashify/launcher/i18n"
	"github.com/MC-Dashify/launcher/utils"
	"github.com/MC-Dashify/launcher/utils/logger"

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

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		logger.Info("Exiting...")
		done <- true
		os.Exit(0)
	}()
	for {
		config.ConfigContent = config.LoadConfig()

		javaFlavour, javaVersion := utils.CheckJava()
		logger.Info(strings.ReplaceAll(strings.ReplaceAll(i18n.Get("java.detected"), "$javaFlavour", javaFlavour), "$javaVersion", javaVersion))

		if strings.HasPrefix(config.ConfigContent.Server, "http") {
			dlServerChannel := make(chan bool)

			go downloadJar([]string{config.ConfigContent.Server}, "server", dlServerChannel)

			dlServerResult := <-dlServerChannel

			if dlServerResult {
				logger.Info("Server file download job is done!")
			} else {
				logger.Info("Server file download job is failed!")
			}
		} else if strings.HasPrefix(config.ConfigContent.Server, "file") {
			logger.Info("Server file source is local, skipping download...")
			serverFilePath = strings.ReplaceAll(config.ConfigContent.Server, "file://", "")
			if _, err := os.Stat(serverFilePath); os.IsNotExist(err) {
				logger.Fatal("Server file not found!")
			}
		} else {
			logger.Fatal("Invalid server file source!")
		}

		if len(config.ConfigContent.Plugins) > 0 {
			dlPluginsChannel := make(chan bool)
			go downloadJar(config.ConfigContent.Plugins, "plugins", dlPluginsChannel)
			dlPluginsResult := <-dlPluginsChannel

			if dlPluginsResult {
				logger.Info("Plugin file(s) download job is done!")
			} else {
				logger.Info("Plugin file(s) download job is failed!")
			}
		} else {
			logger.Info("No plugins to download! Skipping...")
		}

		runtimeArgs := prepareRuntime(runtimeJar{}, config.ConfigContent)

		customArgs := append(append(runtimeArgs.arguments, "-jar"), runtimeArgs.serverFile)

		for _, customArg := range config.ConfigContent.JarArgs {
			customArgs = append(customArgs, customArg)
		}

		utils.RunServer(customArgs)
		if config.ConfigContent.Restart {
			logger.Info("Server will restart in 5 seconds. Press Ctrl+C to cancel")
			fmt.Print("> ")

			select {
			case <-time.After(5000 * time.Millisecond):
				fmt.Print("\n")
				logger.Info("Starting Server...")
			}
		} else {
			logger.Info("Exiting...")
			os.Exit(0)
		}
		if !utils.NormalStatusExit {
			logger.Fatal("There was an error while running server. If you didn't stop the process manually, Try to check 'launcher.conf.json'")
		}
		<-done
	}
}

func downloadJar(urls []string, downloadType string, complete chan<- bool) {
	downloadDest := ""
	results := make(map[string]error)
	dlChannel := make(chan downloadResult)

	if downloadType == "server" {
		currentDir, err := os.Getwd()
		if err != nil {
			logger.Fatal(fmt.Sprintf("Failed to get current working dir: %s", err))
		}
		serverDirectory := ""
		if runtime.GOOS == "windows" {
			serverDirectory = currentDir + "\\.launcher\\jars\\"
		} else {
			serverDirectory = currentDir + "/.launcher/jars/"
		}

		logger.Info(fmt.Sprintf("Checking %s directory...", downloadType))
		utils.CheckFolderExist(serverDirectory)
		downloadDest = serverDirectory

	} else if downloadType == "plugins" {
		currentPath, _ := os.Getwd()
		pluginDirectory := currentPath + "/plugins/"

		logger.Info(fmt.Sprintf("Checking %s directory...", downloadType))
		utils.CheckFolderExist(pluginDirectory)
		downloadDest = pluginDirectory

	} else {
		logger.Fatal("Wrong download type!")
	}

	logger.Info(fmt.Sprintf("Preparing parallel download for %s...", downloadType))
	for _, url := range urls {
		go downloadFile(downloadType, downloadDest, url, dlChannel)
	}

	for i := 0; i < len(urls); i++ {
		downloadResult := <-dlChannel
		results[downloadResult.file] = downloadResult.dlerr
	}

	for downloadedFile, downloadError := range results {
		if downloadError != nil {
			logger.Error(fmt.Sprintf("There was an error while downloading %s: %s", downloadedFile, downloadError))
		}
	}
	logger.Info(fmt.Sprintf("Downloaded all %s files!", downloadType))
	complete <- true
}

func downloadFile(downloadType, downloadDir, url string, err chan<- downloadResult) {
	if downloadType == "server" && !utils.IsValidUrl(url) {
		err <- downloadResult{file: url, dlerr: nil}
		return
	}

	client := grab.NewClient()
	req, _ := grab.NewRequest(downloadDir, url)
	req.NoResume = true
	resp := client.Do(req)

	if _, checkFileErr := os.Stat(resp.Filename); !os.IsNotExist(checkFileErr) {
		logger.Info(fmt.Sprintf("File(%s) is already exist. Skipping download...", resp.Filename))
	} else {
		t := time.NewTicker(time.Second)
		defer t.Stop()
	Loop:
		for {
			select {
			case <-t.C:
				etaTime := time.Until(resp.ETA()).Round(time.Second).String()

				if strings.Contains(etaTime, "-") {
					etaTime = "Calculating..."
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
				logger.Info(fmt.Sprintf("[%s] Downloaded %s of %s | ETA: %s | Download Speed: %s/s", jarPath[len(jarPath)-1], currentDownloaded,
					totalDownloaded,
					etaTime,
					downloadSpeed))

			case <-resp.Done:
				break Loop
			}
		}

		if dlErr := resp.Err(); dlErr != nil {
			logger.Error(fmt.Sprintf("Download failed: %s\n", dlErr))
			err <- downloadResult{file: resp.Filename, dlerr: dlErr}
		}

		jarPath := strings.Split(resp.Filename, "/")

		logger.Info(fmt.Sprintf("[%s] Download complete", jarPath[len(jarPath)-1]))
	}
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
