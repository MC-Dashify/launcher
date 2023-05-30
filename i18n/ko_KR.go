package i18n

var ko_KR map[string]string = map[string]string{
	// 일반
	"general.exec":                                              "명령어 실행: $command",
	"general.exiting":                                           "종료중...",
	"general.calculating":                                       "계산중...",
	"general.server.download.success":                           "서버 파일을 성공적으로 다운로드 하였습니다.",
	"general.plugin.download.success":                           "플러그인 파일(들)을 성공적으로 다운로드 하였습니다.",
	"general.server.download.failed":                            "서버 파일 다운로드에 실패하였습니다.",
	"general.plugin.download.failed":                            "플러그인 파일 다운로드 실패하였습니다.",
	"general.server.source.local":                               "서버 파일 소스가 로컬임을 감지했습니다. 다운로드를 건너뜁니다.",
	"general.server.source.local.notfound.or.permission.denied": "서버 파일 소스가 로컬임을 감지했지만, 해당 경로에 파일이 없거나 접근 권한이 없습니다. 구성 설정 파일을 확인해주세요.",
	"general.server.source.invalid.protocol":                    "서버 파일 소스 프로토콜이 올바르지 않습니다. 구성 설정 파일을 확인해주세요.",
	"general.server.starting":                                   "서버 시작중...",
	"general.server.crashed":                                    "서버가 충돌한 것 같습니다. 수동으로 서버 프로세스를 강제 종료 한 것이 아니라면, 구성 설정 파일이 올바른지 확인해주세요.",
	"general.plugin.empty":                                      "다운로드할 플러그인이 없습니다. 건너뛰는중...",
	"general.server.restart":                                    "서버가 5초 뒤에 재시작 됩니다. Ctrl + C를 누르거나 이 창 또는 터미널을 닫아서 종료할 수 있습니다.",
	"general.cwd.get.failed":                                    "현재 작업 경로를 알 수 없습니다. 오류 정보: $error",
	"general.checking.directory":                                "$dir를 위한 경로 확인중...",
	"general.download.type.invalid":                             "올바르지 않은 다운로드 유형 발견",
	"general.download.preparing":                                "$file을(를) 위한 병렬 다운로드 준비중...",
	"general.download.failed":                                   "파일 $file을(를) 다운로드 받는데 실패하였습니다. 오류 정보: $error",
	"general.download.success":                                  "성공적으로 모든 $type 파일(들)을 다운로드 하였습니다.",
	"general.download.file.exist":                               "파일($file)이 이미 존재합니다. 다운로드를 건너뛰는중...",
	"general.download.progress":                                 "[$fileName] 전체 $fileSize중 $downloadedSize 다운로드 됨 | 예상 남은 시간: $eta | 다운로드 속도: $downloadSpeed/s",
	"general.download.done":                                     "[$fileName] 다운로드 완료.",
	"general.unsafe.shutdown":                                   "안전하지 않은 종료가 감지되었습니다. 'stop' 명령어를 사용하여 서버를 정상적으로 종료해 주세요. Ctrl + C 또는 Ctrl + D와 같은 강제 종료 명령어를 사용하지 않는 것을 권장합니다.",

	// Java
	"java.detected":       "$javaVersion 버전의 $javaFlavour Java를 감지했습니다.",
	"java.notfound":       "Java를 찾을 수 없습니다. Java를 설치하거나 PATH 환경 변수를 확인해주세요.",
	"java.jvm.stopped":    "JVM Runtime 종료됨.",
	"java.jvm.fail.start": "JVM Runtime을 시작하는데 실패했습니다. 오류 정보: $error",

	// 인자
	"flag.lang.desc":        "표시할 언어를 선택합니다. 인자는 'en-US' 또는 'ko-KR' 같은 형식이어야 합니다.",
	"flag.verbose.desc":     "상세한 로그를 포함하여 출력합니다.",
	"flag.version.desc":     "MC-Dashify 실행기 버전을 보여줍니다.",
	"flag.config.help.desc": "MC-Dashify 구성 설정 파일에 대한 도움말을 보여줍니다.",

	// 파일
	"file.generating.missings": "찾을 수 없는 폴더(들) 생성중...",
	"file.unknown.size":        "알 수 없는 크기",
	"file.info.fetch.failed":   "파일 정보를 불러오는데 실패하였습니다. 오류 정보: $error",

	// 네트워크
	"net.file.info.fetch.failed":      "URL에서 파일 정보를 불러오는데 실패하였습니다. 오류 정보: $error",
	"net.file.info.time.fetch.failed": "URL에서 파일의 마지막 수정 날짜를 불러오는데 실패하였습니다. 오류 정보: $error",

	// 버전
	"version.invalid": "버전 $version은 올바르지 않습니다.",
	"version.info":    "MC-Dashify launcher v.$version",

	// WebConsole
	"webconsole.started1":                       "+-------------------------+",
	"webconsole.started2":                       "| WebConsole 서버 시작됨! |",
	"webconsole.chk.valid.prev.connection":      "[WebConsole] 유효한 이전 연결 확인중...",
	"webconsole.restoring.prev.connection":      "[WebConsole] 이전 연결 복원중: $connection",
	"webconsole.connection.closed":              "[WebConsole] 연결 종료됨: $connection",
	"webconsole.connection.closed.error":        "[WebConsole] 다음 오류로 인해 $remote와의 연결이 종료되었습니다: $error",
	"webconsole.connection.close.msg.send.fail": "[WebConsole] $remote에 연결 종료 메세지를 보내는데 실패했습니다. 오류 정보: $error",
	"webconsole.connection.opened":              "[WebConsole] 연결 수립됨: $connection",
	"webconsole.connection.cmd.received":        "[WebConsole] $remote로부터의 명령: $command",

	// 구성 설정 관련
	"config.notfound":                "구성 설정 파일을 찾을 수 없습니다. 새로 생성중...",
	"config.empty":                   "구성 설정 파일이 비어있습니다. 새로 생성중...",
	"config.invalid":                 "구성 설정이 올바르지 않습니다. 구성 설정 파일을 확인해주세요. 자세한 오류 내용은 다음과 같습니다: $error",
	"config.create_failed":           "구성 설정 파일을 생성할 수 없습니다. 권한을 확인해주세요. 자세한 오류 내용은 다음과 같습니다: $error",
	"config.write_failed":            "구성 설정 파일을 작성할 수 없습니다. 권한을 확인해주세요. 자세한 오류 내용은 다음과 같습니다: $error",
	"config.created":                 "구성 설정 파일을 성공적으로 생성하였습니다.",
	"config.server.empty":            "서버 파일의 경로나 URL이 비어있습니다.",
	"config.memory.invalid":          "올바르지 않은 메모리 설정을 발견하였습니다. 메모리는 GB 단위로 설정되며, 최소 2GB 이상이 필요합니다. 구성 설정 파일을 확인해주세요.",
	"config.debug_port.invalid":      "올바르지 않은 JVM 디버깅 포트 설정을 발견하였습니다. 구성 설정 파일을 확인해주세요.",
	"config.api_port.invalid":        "올바르지 않은 API 포트 설정을 발견하였습니다. 구성 설정 파일을 확인해주세요.",
	"config.plugin_api_port.invalid": "올바르지 않은 플러그인 API 포트 설정을 발견하였습니다. 구성 설정 파일을 확인해주세요.",
	"config.help": "MC-Dashify 실행기 구성 설정 파일 도움말입니다.\n\n" +
		"- \"server\" 항목에는 \"http://\" 또는 \"https://\"로 시작하는 서버 파일의 URL 또는 \"file://\"로 시작하는 로컬 경로를 입력해야 합니다.\n" +
		"  만약 file://을 이용하여 로컬 파일을 이용하여 실행 할 경우, 전체 경로를 입력해야 합니다.\n\n" +
		"- \"debug\" 항목에는 개발을 위한 JVM 디버그 모드를 활성화 여부를 \"bool값으로\" 입력해야 합니다.\n\n" +
		"- \"debug_port\" 항목에는 개발을 위한 JVM 디버그 모드에 사용될 포트를 입력해야 합니다.\n" +
		"  \"debug\" 항목이 false로 설정되어 있다면 이 옵션은 무시됩니다.\n\n" +
		"- \"api_port\" 항목에는 서버가 통신할 API 서버의 포트를 입력해야 합니다.\n\n" +
		"- \"plugin_api_port\" 항목에는 MC-Dashify 플러그인과 통신할 API 서버의 포트를 입력해야 합니다.\n\n" +
		"- 모든 \"포트 관련 항목\"은 서로 같은 값을 사용할 수 없습니다.\n\n" +
		"- \"memory\" 항목에는 Minecraft 서버에 할당될 메모리의 양을 입력해야 합니다.\n  GB 단위로 입력해야 하며, 최소 2GB 이상이 필요합니다.\n\n" +
		"- \"plugins\" 항목에는 서버에 같이 들어갈 플러그인의 URL 주소를 입력할 수 있습니다.\n  별도로 다운받고 싶지 않다면 비워둘 수 있습니다.\n\n" +
		"- \"jarArgs\" 항목에는 Minecraft 서버에 들어갈 인수를 설정할 수 있습니다.\n  여기서 인수는 \"nogui\" 같은 인수를 뜻합니다.\n  인수는 비워둘 수 있습니다.",
}
