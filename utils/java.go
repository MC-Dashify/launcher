package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/MC-Dashify/launcher/i18n"
	"github.com/MC-Dashify/launcher/utils/logger"
)

var NormalStatusExit bool = false

func CheckJava() (javaFlavor, javaVersion string) {
	out, err := staticExecutor("java", []string{"-version"})
	if err != nil {
		logger.Error(i18n.Get("java.notfound"))
		os.Exit(1)
	}
	javaFlavor = strings.ReplaceAll(strings.Split(out, " ")[0], "\"", "")
	javaVersion = strings.ReplaceAll(strings.Split(out, " ")[2], "\"", "")
	return javaFlavor, javaVersion
}

func RunServer(arguments []string) {
	interactiveExecutor("java", arguments)
}

func SelectOptionByMemory(memory int) []string {
	memoryOptions := []string{"--add-modules=jdk.incubator.vector"}
	if memory >= 12 {
		logger.Info("Using Aikar's Advanced memory options")
		for _, option := range []string{
			"-XX:G1NewSizePercent=40",
			"-XX:G1MaxNewSizePercent=50",
			"-XX:G1HeapRegionSize=16M",
			"-XX:G1ReservePercent=15",
			"-XX:InitiatingHeapOccupancyPercent=20",
		} {
			memoryOptions = append(memoryOptions, option)
		}
	} else {
		logger.Info("Using Aikar's standard memory options")
		for _, option := range []string{
			"-XX:G1NewSizePercent=30",
			"-XX:G1MaxNewSizePercent=40",
			"-XX:G1HeapRegionSize=8M",
			"-XX:G1ReservePercent=20",
			"-XX:InitiatingHeapOccupancyPercent=15",
		} {
			memoryOptions = append(memoryOptions, option)
		}
	}
	return memoryOptions
}

func staticExecutor(baseCmd string, cmdArgs []string) (string, error) {
	logger.Debug(fmt.Sprintf("Exec: %v", baseCmd+" "+strings.Join(cmdArgs, " ")))

	cmd := exec.Command(baseCmd, cmdArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func interactiveExecutor(baseCmd string, cmdArgs []string) error {
	logger.Debug(fmt.Sprintf("Exec: %v", baseCmd+" "+strings.Join(cmdArgs, " ")))

	cmd := exec.Command(baseCmd, cmdArgs...)
	env := os.Environ()
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // setting this allowed me to interact with ncurses interface from `make menuconfig`

	err := cmd.Start()
	if err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	NormalStatusExit = true
	return nil
}
