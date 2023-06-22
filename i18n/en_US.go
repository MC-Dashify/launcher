package i18n

var en_US map[string]string = map[string]string{
	// General
	"general.exec":                                              "Executing command: $command",
	"general.exiting":                                           "Exiting...",
	"general.calculating":                                       "Calculating...",
	"general.server.download.success":                           "Successfully downloaded server file.",
	"general.plugin.download.success":                           "Successfully downloaded plugin file(s).",
	"general.server.download.failed":                            "The server file download failed.",
	"general.plugin.download.failed":                            "The plugin file(s) download failed.",
	"general.server.source.local":                               "The server file source is local. Skipping download...",
	"general.server.source.local.notfound.or.permission.denied": "The server file source is local, but it seems the file does not exist on a specified path, or permission denied. Please check and try again.",
	"general.server.source.invalid.protocol":                    "The server file source protocol is not valid. Please check your config file.",
	"general.server.starting":                                   "Starting server...",
	"general.server.crashed":                                    "It seems the server crashed. If you didn't stop or killed the process manually, Try to check 'launcher.conf.json'",
	"general.plugin.empty":                                      "No plugins to download! Skipping download...",
	"general.server.restart":                                    "The server will restart in 5 seconds. You can exit by pressing Ctrl + C or closing this window or terminal.",
	"general.cwd.get.failed":                                    "Failed to get current working directory. Error detail: $error",
	"general.checking.directory":                                "Checking $dir directory...",
	"general.download.type.invalid":                             "Wrong download type!",
	"general.download.preparing":                                "Preparing parallel download for $file...",
	"general.download.failed":                                   "Failed to download $file. Error detail: $error",
	"general.download.success":                                  "Successfully downloaded all $type file(s).",
	"general.download.file.exist":                               "File($file) already exists. Skipping download...",
	"general.download.progress":                                 "[$fileName] Downloaded $downloadedSize of $fileSize | ETA: $eta | Download Speed: $downloadSpeed/s",
	"general.download.done":                                     "[$fileName] Download Complete.",
	"general.unsafe.shutdown":                                   "Unsafe process kill detected. Please use 'stop' command to stop the server. Do not use Ctrl + C or Ctrl + D to stop the server.",

	//Java
	"java.detected":       "Detected $javaFlavour $javaVersion.",
	"java.notfound":       "Java not found. Please install Java or check your PATH environment variable.",
	"java.jvm.stopped":    "JVM Runtime stopped.",
	"java.jvm.fail.start": "Failed to start JVM Runtime. Error detail: $error",

	// Flags
	"flag.lang.desc":        "Language to disply. Value should be like 'en-US' or 'ko-KR'.",
	"flag.verbose.desc":     "Prints all verbose logs.",
	"flag.mcorigin.desc":    "Change Minecraft listening port. This option is valid only when traffic monitoring is enabled in config file.",
	"flag.version.desc":     "Shows MC-Dashify launhcer version.",
	"flag.config.help.desc": "Shows help about MC-Dashify config file.",

	// Files
	"file.generating.missings": "Generating missing folders...",
	"file.unknown.size":        "Unknown size",
	"file.info.fetch.failed":   "Failed to fetch file info. Error detail: $error",

	// Net
	"net.file.info.fetch.failed":      "Failed to fetch file info from url. Error detail: $error",
	"net.file.info.time.fetch.failed": "Failed to parse last modified time of file from url. Error detail: $error",
	"traffic.monitor.enabled":         "[TrafficMonitor] Traffic monitoring is enabled. To measure traffic, connect to the $redirectPort port. If you changed the Minecraft server port, you must change the Minecraft server port using the --mcorigin flag. (ex: --mcorigin 25565)",

	// Version
	"version.invalid": "Version $version is invalid.",
	"version.info":    "MC-Dashify launcher v.$version",

	// WebConsole
	"webconsole.started1":                       "+----------------------------+",
	"webconsole.started2":                       "| WebConsole Server Started! |",
	"webconsole.chk.valid.prev.connection":      "[WebConsole] Checking Valid Previous Connections...",
	"webconsole.restoring.prev.connection":      "[WebConsole] Restoring Previous Connection: $connection",
	"webconsole.connection.closed":              "[WebConsole] Connection Closed: $connection",
	"webconsole.connection.closed.error":        "[WebConsole] Connection from $remote closed due to following error: $error",
	"webconsole.connection.close.msg.send.fail": "[WebConsole] Failed to send connection close message to $remote. Error detail: $error",
	"webconsole.connection.opened":              "[WebConsole] Connection Opened: $connection",
	"webconsole.connection.cmd.received":        "[WebConsole] FROM $remote CMD: $command",

	// Config
	"config.notfound":                "Config file not found. Creating new one...",
	"config.empty":                   "Config file is empty. Creating new one...",
	"config.invalid":                 "Config file is invalid. Please check your config file. Error detail: $error",
	"config.create_failed":           "Failed to create config file. Please check your permission. Error detail: $error",
	"config.write_failed":            "Failed to write config file. Please check your permission. Error detail: $error",
	"config.created":                 "Config file created successfully.",
	"config.version.different":       "Config file version is different. Consider updating your config file.",
	"config.server.empty":            "Server file path or URL is empty.",
	"config.memory.invalid":          "Invalid memory option detected. Memory is set in GB and requires at least 2 GB. Please check the config file.",
	"config.debug_port.invalid":      "Invalid JVM debugging port settings found. Please check the config file.",
	"config.api_port.invalid":        "Invalid API port. Please check your config file.",
	"config.plugin_api_port.invalid": "Invalid Plugin API port. Please check your config file.",
}
