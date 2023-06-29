<p align="center">
  <img width="128" align="center" src="https://github.com/MC-Dashify/launcher/blob/main/.github/assets/logo-512.png">
</p>
<h1 align="center">Dashify</h1>
<h3 align="center">Easily and quickly monitor Minecraft servers</h3>
<p align="center">Standalone monitoring system without any additional installation</p>
<p align="center">
  <a href="https://github.com/MC-Dashify/launcher/actions/workflows/codeql.yml">
    <img src="https://github.com/MC-Dashify/launcher/actions/workflows/codeql.yml/badge.svg" alt="CodeQL" />
  </a>
  <a href="https://github.com/MC-Dashify/launcher/actions/workflows/main.yml">
    <img src="https://github.com/MC-Dashify/launcher/actions/workflows/main.yml/badge.svg" alt="Build & Publish to Release" />
  </a>
  <a href="https://app.codacy.com/gh/MC-Dashify/launcher/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade"><img src="https://app.codacy.com/project/badge/Grade/bf744aecc0c44413914e2fcecda263f0" alt="code quality"/></a>
</p>

<p align="center"><a href="https://github.com/MC-Dashify/launcher/blob/main/README.md">English</a> · <a href="https://github.com/MC-Dashify/launcher/blob/main/.github/documents/README.ko_KR.md">한국어</a></p>

<h1 align="center">THIS IS LAUNCHER REPOSITORY FOR DASHIFY.</h1>

## Configuration

Please refer to [this](https://github.com/MC-Dashify/launcher/blob/main/.github/documents/CONFIG_GUIDE.md).

## API Endpoint

- NOTE: All HTTP requests requires `Authorization` header like `Bearer your_key_here`.
  > Key information can be found in the file `Dashify/config.yml` in the server plugin folder.

<table>
<thead>
  <tr>
    <th>Path</th>
    <th>Description</th>
    <th>Protocol</th>
    <th>Request Method</th>
    <th>Required Params</th>
    <th>Request Example</th>
    <th>Note</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td><code>/console</code></td>
    <td>Connect to Console's I/O Directly</td>
    <td><code>Web Socket</code></td>
    <td><code style="color:#ff8888">N/A</code></td>
    <td><code>auth_key [string]</code></td>
    <td><code>ws://localhost:8080/console?auth_key=your_key_here</code></td>
    <td><code style="color:#ff8888">N/A</code></td>
  </tr>
  <tr>
    <td><code>/ping</code></td>
    <td>Ping the server</td>
    <td><code>HTTP</code></td>
    <td><code style="color:#6bdd9a">GET</code></td>
    <td><code style="color:#ff8888">N/A</code></td>
    <td><code>http://localhost:8080/ping</code></td>
    <td><code style="color:#ff8888">N/A</code></td>
  </tr>
  <tr>
    <td><code>/logs</code></td>
    <td>Get Minecraft server's logs</td>
    <td><code>HTTP</code></td>
    <td><code style="color:#6bdd9a">GET</code></td>
    <td><code>lines [int]</code></td>
    <td><code>http://localhost:8080/logs?lines=100</code></td>
    <td>Parameter <code style="color:#cc00cc">line</code> should a valid int between <code>1</code> and <code>1000</code></td>
  </tr>
  <tr>
    <td><code>/traffic</code></td>
    <td>Get Minecraft server's traffic information</td>
    <td><code>HTTP</code></td>
    <td><code style="color:#6bdd9a">GET</code></td>
    <td><code style="color:#ff8888">N/A</code></td>
    <td><code>http://localhost:8080/traffic</code></td>
    <td>Server keeps count traffic information until request have received. After respond, traffic will be reset (Recount after request)</td>
  </tr>
</tbody>
</table>

Other endpoints are ReverseProxyed. They are documented at [plugin](https://github.com/MC-Dashify/plugin) repository. Just remember that port number have to same with launcher's port number.

## Code of Conduct

See the [CODE_OF_CONDUCT.md](https://github.com/MC-Dashify/launcher/blob/main/CODE_OF_CONDUCT.md) file for Code of Conduct information.

## Contributing

See the [CONTRIBUTING.md](https://github.com/MC-Dashify/launcher/blob/main/CONTRIBUTING.md) file for contributing information.

## License

See the [LICENSE](https://github.com/MC-Dashify/launcher/blob/main/LICENSE) file for licensing information.
