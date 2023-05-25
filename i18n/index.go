package i18n

import (
	"fmt"

	"github.com/MC-Dashify/launcher/global"
	"github.com/MC-Dashify/launcher/utils/logger"
	"github.com/Xuanwo/go-locale"
)

func Get(key string) string {
	language := "en-US"
	if global.IsLanguageForced {
		language = global.ForcedLanguage
	} else {
		tag, err := locale.Detect()
		if err != nil {
			logger.Fatal(fmt.Sprintf("Failed to detect locale: %v", err.Error()))
		}
		language = tag.String()
	}

	// tag.String() returns the language code and the country code, e.g. "en-US"
	switch language {
	case "ko-KR":
		return ko_KR[key]
	default:
		return en_US[key]
	}
}
