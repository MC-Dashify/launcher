package utils

import (
	"os"
	"os/exec"
	"strings"

	"github.com/MC-Dashify/launcher/i18n"
	"github.com/MC-Dashify/launcher/utils/logger"
)

func CheckJava() (javaFlavor, javaVersion string) {
	out, err := StaticExecutor("java", []string{"-version"})
	if err != nil {
		logger.Error(i18n.Get("java.notfound"))
		os.Exit(1)
	}
	javaFlavor = strings.ReplaceAll(strings.Split(out, " ")[0], "\"", "")
	javaVersion = strings.ReplaceAll(strings.Split(out, " ")[2], "\"", "")
	return javaFlavor, javaVersion
}

func SelectOptionByMemory(memory int) []string {
	memoryOptions := []string{"-Dfile.encoding=UTF-8", "--add-modules=jdk.incubator.vector"}
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

func StaticExecutor(baseCmd string, cmdArgs []string) (string, error) {
	logger.Debug(strings.ReplaceAll(i18n.Get("general.exec"), "$command", baseCmd+" "+strings.Join(cmdArgs, " ")))

	cmd := exec.Command(baseCmd, cmdArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
