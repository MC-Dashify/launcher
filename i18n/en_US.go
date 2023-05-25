package i18n

var en_US map[string]string = map[string]string{
	"test.hello":        "Hello",
	"test.placeholder":  "Hello $user!",
	"flag.lang.desc":    "Language to disply. Value should be like 'en-US' or 'ko-KR'.",
	"flag.verbose.desc": "Prints all verbose logs.",
	"java.detected":     "Detected $javaFlavour $javaVersion.",
	"java.notfound":     "Java not found. Please install Java or check your PATH environment variable.",
}
