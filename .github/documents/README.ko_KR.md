<p align="center">
  <img width="128" align="center" src="https://github.com/MC-Dashify/launcher/blob/main/.github/assets/logo-512.png">
</p>
<h1 align="center">Dashify</h1>
<h3 align="center">Minecraft 서버 모니터링을 쉽고 빠르게</h3>
<p align="center">별도의 설치 없이 실행되는 Standalone 모니터링 시스템</p>
<p align="center">
  <a href="https://github.com/MC-Dashify/launcher/actions/workflows/main.yml">
    <img src="https://github.com/MC-Dashify/launcher/actions/workflows/main.yml/badge.svg" alt="Build & Publish to Release" />
  </a>
</p>

<p align="center"><a href="https://github.com/MC-Dashify/launcher/blob/main/README.md">English</a> · <a href="https://github.com/MC-Dashify/launcher/blob/main/.github/documents/README.ko_KR.md">한국어</a></p>

<h1 align="center">이 레포지토리는 Dashify의 launcher 레포지토리 입니다.</h1>

## 구성 설정

[여기](https://github.com/MC-Dashify/launcher/blob/main/.github/documents/CONFIG_GUIDE.ko_KR.md)를 참고해주세요.

## API Endpoint

- 참고: 모든 HTTP 요청에는 `Authorization` 헤더에 `Bearer your_key_here` 와 같은 형식의 토큰 인증이 필요합니다.
  > 키 정보는 서버 플러그인 폴더 내 `Dashify/config.yml` 파일에서 확인할 수 있습니다.

<table>
<thead>
  <tr>
    <th>경로</th>
    <th>요약</th>
    <th>프로토콜</th>
    <th>요청 메서드</th>
    <th>필요한 파라미터</th>
    <th>요청 예시</th>
    <th>비고</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td><code>/console</code></td>
    <td>콘솔 I/O에 직접적으로 접근합니다.</td>
    <td><code>Web Socket</code></td>
    <td><code style="color:#ff8888">N/A</code></td>
    <td><code>auth_key [string]</code></td>
    <td><code>ws://localhost:8080/console?auth_key=your_key_here</code></td>
    <td><code style="color:#ff8888">N/A</code></td>
  </tr>
  <tr>
    <td><code>/ping</code></td>
    <td>서버가 작동 중인지 확인합니다.</td>
    <td><code>HTTP</code></td>
    <td><code style="color:#6bdd9a">GET</code></td>
    <td><code style="color:#ff8888">N/A</code></td>
    <td><code>http://localhost:8080/ping</code></td>
    <td><code style="color:#ff8888">N/A</code></td>
  </tr>
  <tr>
    <td><code>/logs</code></td>
    <td>Minecraft 서버의 로그를 가져옵니다.</td>
    <td><code>HTTP</code></td>
    <td><code style="color:#6bdd9a">GET</code></td>
    <td><code>lines [int]</code></td>
    <td><code>http://localhost:8080/logs?lines=100</code></td>
    <td>파라미터 <code style="color:#cc00cc">line</code>은 반드시<code>1</code>부터 <code>1000</code> 사이의 유효한 정수이어야 합니다.</td>
  </tr>
  <tr>
    <td><code>/traffic</code></td>
    <td>Minecraft 서버의 트래픽 정보를 가져옵니다.</td>
    <td><code>HTTP</code></td>
    <td><code style="color:#6bdd9a">GET</code></td>
    <td><code style="color:#ff8888">N/A</code></td>
    <td><code>http://localhost:8080/traffic</code></td>
    <td>서버는 요청이 올 때 까지 트래픽을 누산합니다. 요청이 오면 현재까지 누산된 트래픽을 반환하고 리셋됩니다.</td>
  </tr>
</tbody>
</table>

다른 경로들은 역방향 프록시 처리 되어 있습니다. 해당 경로들은 [plugin](https://github.com/MC-Dashify/plugin) 레포지토리에서 확인할 수 있습니다. launcher의 포트번호를 사용해야 한 다는 것을 주의하세요.

## Code of Conduct

Code of Conduct 파일을 보거나 내용을 확인하려면 [CODE_OF_CONDUCT.md](https://github.com/MC-Dashify/launcher/blob/main/.github/documents/CODE_OF_CONDUCT.ko_KR.md) 파일을 확인하세요.

## Contributing

기여 가이드라인을 확인하려면 [CONTRIBUTING.md](https://github.com/MC-Dashify/launcher/blob/main/.github/documents/CONTRIBUTING.ko_KR.md) 파일을 확인하세요.

## License

라이센스를 확인하려면 [LICENSE](https://github.com/MC-Dashify/launcher/blob/main/LICENSE) 파일을 확인하세요.
