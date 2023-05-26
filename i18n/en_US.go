package i18n

var en_US map[string]string = map[string]string{
	"test.hello":                     "Hello",
	"test.placeholder":               "Hello $user!",
	"flag.lang.desc":                 "Language to disply. Value should be like 'en-US' or 'ko-KR'.",
	"flag.verbose.desc":              "Prints all verbose logs.",
	"java.detected":                  "Detected $javaFlavour $javaVersion.",
	"java.notfound":                  "Java not found. Please install Java or check your PATH environment variable.",
	"config.notfound":                "Config file not found. Creating new one...",
	"config.empty":                   "Config file is empty. Creating new one...",
	"config.invalid":                 "Config file is invalid. Please check your config file. Error detail: $error",
	"config.create_failed":           "Failed to create config file. Please check your permission. Error detail: $error",
	"config.write_failed":            "Failed to write config file. Please check your permission. Error detail: $error",
	"config.created":                 "Config file created successfully.",
	"config.server.empty":            "Server file path or URL is empty.",
	"config.api_port.invalid":        "Invalid API port. Please check your config file.",
	"config.plugin_api_port.invalid": "Invalid Plugin API port. Please check your config file.",
}
