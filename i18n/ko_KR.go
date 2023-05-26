package i18n

var ko_KR map[string]string = map[string]string{
	"test.hello":                     "안녕하세요",
	"test.placeholder":               "안녕하세요 $user 님!",
	"flag.lang.desc":                 "표시할 언어를 선택합니다. 인자는 'en-US' 또는 'ko-KR' 같은 형식이어야 합니다.",
	"flag.verbose.desc":              "상세한 로그를 포함하여 출력합니다.",
	"java.detected":                  "$javaVersion 버전의 $javaFlavour Java를 감지했습니다.",
	"java.notfound":                  "Java를 찾을 수 없습니다. Java를 설치하거나 PATH 환경 변수를 확인해주세요.",
	"config.notfound":                "구성 설정 파일을 찾을 수 없습니다. 새로 생성중...",
	"config.empty":                   "구성 설정 파일이 비어있습니다. 새로 생성중...",
	"config.invalid":                 "구성 설정이 올바르지 않습니다. 구성 설정 파일을 확인해주세요. 자세한 오류 내용은 다음과 같습니다: $error",
	"config.create_failed":           "구성 설정 파일을 생성할 수 없습니다. 권한을 확인해주세요. 자세한 오류 내용은 다음과 같습니다: $error",
	"config.write_failed":            "구성 설정 파일을 작성할 수 없습니다. 권한을 확인해주세요. 자세한 오류 내용은 다음과 같습니다: $error",
	"config.created":                 "구성 설정 파일을 성공적으로 생성하였습니다.",
	"config.server.empty":            "서버 파일의 경로나 URL이 비어있습니다.",
	"config.api_port.invalid":        "올바르지 않은 API 포트 설정을 발견하였습니다. 구성 설정 파일을 확인해주세요.",
	"config.plugin_api_port.invalid": "올바르지 않은 플러그인 API 포트 설정을 발견하였습니다. 구성 설정 파일을 확인해주세요.",
}
