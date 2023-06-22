# 구성 설정 파일 가이드

<p align="center"><a href="https://github.com/MC-Dashify/launcher/blob/main/.github/documents/CONFIG_GUIDE.md">English</a> · <a href="https://github.com/MC-Dashify/launcher/blob/main/.github/documents/CONFIG_GUIDE.ko_KR.md">한국어</a></p>

## 구성 설정별 설명

구성 설정 파일은 JSON 형식으로 작성되어 있으며, `launcher.config.json`이라는 이름으로 자동 생성됩니다.

<table>
<thead>
  <tr>
    <th>키</th>
    <th>설명</th>
    <th>기본값</th>
    <th>비고</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td><code>config_version</code></td>
    <td>현재 구성 설정의 버전을 나타냅니다.</td>
    <td><code>1</code></td>
    <td><span style="color:#ff8888">구성 설정 버전을 임의로 변경할 경우 의도치 않은 오류가 일어날 수 있습니다.</span></td>
  </tr>
   <tr>
    <td><code>server</code></td>
    <td>서버 jar 파일의 경로 또는 링크를 설정합니다.</td>
    <td><code>https://clip.aroxu.me/download?mc_version=1.19.4</code></td>
    <td><span>http(s)://로 시작하는 서버 다운로드 링크를 넣거나 file://path/to/server.jar 형식의 서버 파일 위치를 넣습니다.</span></td>
  </tr>
  <tr>
    <td><code>debug</code></td>
    <td>JVM 디버깅 활성화 여부를 설정합니다.</td>
    <td><code>false</code></td>
    <td><span>N/A</span></td>
  </tr>
  <tr>
    <td><code>debug_port</code></td>
    <td>JVM 디버깅 포트를 설정합니다.</td>
    <td><code>5005</code></td>
    <td><span style="color:#ff8888"><code>debug</code> 옵션이 꺼져있을 경우 이 옵션은 무시됩니다.</span></td>
  </tr>
  <tr>
    <td><code>restart</code></td>
    <td>서버 JVM 런타임이 종료되었을 경우 자동으로 재시작 할지 설정합니다.</td>
    <td><code>true</code></td>
    <td><span style="color:#ff8888">불가항적 비정상 종료(정전, 커널 패닉 또는 블루스크린)에 대해서는 해당되지 않습니다.</span></td>
  </tr>
  <tr>
    <td><code>memory</code></td>
    <td>서버 JVM 런타임에 할당할 메모리 양을 GB단위로 설정합니다.</td>
    <td><code>2</code></td>
    <td><span style="color:#ff8888">2GB가 넘지 않을경우 서버가 실행 되지 않을 수 있습니다.</span></td>
  </tr>
  <tr>
    <td><code>enable_traffic_monitor</code></td>
    <td>서버 트래픽 측정 기능을 활성화 할지 설정합니다.</td>
    <td><code>false</code></td>
    <td><span style="color:#ff8888">해당 기능은 아직 alpha 기능입니다. README.MD 파일을 참고하여 알려진 이슈 목록을 확인하세요.</span></td>
  </tr>
  <tr>
    <td><code>enable_traffic_monitor</code></td>
    <td>서버 트래픽 측정 기능을 사용하려면 포트를 리다이렉트 해야 합니다. 리다이렉트 할 때 사용할 포트를 설정합니다.</td>
    <td><code>25555</code></td>
    <td><span style="color:#ff8888">해당 기능은 아직 alpha 기능입니다. README.MD 파일을 참고하여 알려진 이슈 목록을 확인하세요. <code>enable_traffic_monitor</code> 옵션이 꺼져있을 경우 이 옵션은 무시됩니다.</span></td>
  </tr>
  <tr>
    <td><code>api_port</code></td>
    <td>밖으로 노출되는 포트를 설정합니다.</td>
    <td><code>8080</code></td>
    <td><span>N/A</span></td>
  </tr>
  <tr>
    <td><code>plugin_api_port</code></td>
    <td>서버에 들어가는 Dashify 플러그인 설정에 들어가있는 포트와 동일하게 해주세요. Dashify는 해당 포트를 통해 정보를 불러옵니다. 포트가 일치하지 않을 경우 제대로 된 정보가 전달되지 않습니다.</td>
    <td><code>8081</code></td>
    <td><span>플러그인 설정 파일은 `plugins/Dashify/config.yml` 경로에 있습니다.</span></td>
  </tr>
  <tr>
    <td><code>plugins</code></td>
    <td>서버에 들어갈 플러그인의 다운로드 링크를 배열 형태로 적어주세요.</td>
    <td><code>["https://github.com/MC-Dashify/plugin/releases/latest/download/dashify-plugin-all.jar"]</code></td>
    <td><span>링크가 올바르지 않을 경우 해당 링크는 무시됩니다. 여기에 적힌 플러그인들은 launcher가 시작될 때 마다 최신 버전으로 다운로드 됩니다.</span></td>
  </tr>
  <tr>
    <td><code>jar_args</code></td>
    <td>서버 실행에 들어갈 arguments를 설정합니다.</td>
    <td><code>["nogui"]</code></td>
    <td><span style="color:#ff8888">유효하지 않은 argument는 무시됩니다. 일부 옵션을 잘못 입력하면 서버가 제대로 작동하지 않을 수 있습니다.</span></td>
  </tr>
</tbody>
</table>
