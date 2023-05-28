package global

import (
	"os/exec"
)

const Version = "0.0.1"

var IsVerbose bool = false

var IsLanguageForced bool = false
var ForcedLanguage string = ""

var Cmd *exec.Cmd

var JarArgs []string

var NormalStatusExit bool = false
var IsMCServerRunning bool = false
