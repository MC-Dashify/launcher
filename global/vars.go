package global

import (
	"os/exec"
	"sync"
)

var IsVerbose bool = false

var IsLanguageForced bool = false
var ForcedLanguage string = ""

var Cmd *exec.Cmd

var JarArgs []string

var NormalStatusExit bool = true
var IsMCServerRunning bool = false

var MCOriginPort int = 25565

type TrafficClientStats struct {
	ReceivedBytes int64
	SentBytes     int64
}

var TrafficClients = make(map[string]*TrafficClientStats)
var TrafficClientsMutex = sync.RWMutex{}
