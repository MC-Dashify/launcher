# Config File Guide

<p align="center"><a href="https://github.com/MC-Dashify/launcher/blob/main/.github/documents/CONFIG_GUIDE.md">English</a> · <a href="https://github.com/MC-Dashify/launcher/blob/main/.github/documents/CONFIG_GUIDE.ko_KR.md">한국어</a></p>

## Description by Configuration Settings

The configuration settings file is created in JSON format and is automatically generated with the name `launcher.config.json`.

<table>
<thead>
  <tr>
    <th>Key</th>
    <th>Description</th>
    <th>Default Value</th>
    <th>Note</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td><code>config_version</code></td>
    <td>Indicates the version of the current configuration settings.</td>
    <td><code>1</code></td>
    <td><span style="color:#ff8888">Random changes to the configuration settings version can result in unintended errors.</span></td>
  </tr>
   <tr>
    <td><code>server</code></td>
    <td>Sets the path or link to the server jar file.</td>
    <td><code>https://clip.aroxu.me/download?mc_version=1.19.4</code></td>
    <td><span>Insert the server download link that starts with http(s):// or enter the server file location like following format: file://path/to/server.jar.</span></td>
  </tr>
  <tr>
    <td><code>debug</code></td>
    <td>Set whether to enable JVM debugging.</td>
    <td><code>false</code></td>
    <td><span>N/A</span></td>
  </tr>
  <tr>
    <td><code>debug_port</code></td>
    <td>Sets the JVM debugging port.</td>
    <td><code>5005</code></td>
    <td><span style="color:#ff8888">If the <code>debug</code> option is turned off, this option is ignored.</span></td>
  </tr>
  <tr>
    <td><code>restart</code></td>
    <td>Set whether to restart automatically when the server JVM runtime ends.</td>
    <td><code>true</code></td>
    <td><span style="color:#ff8888">Not applicable for force majeure abnormal termination (power outage, kernel panic, or blue screen).</span></td>
  </tr>
  <tr>
    <td><code>memory</code></td>
    <td>Sets the amount of memory in GB to allocate at server JVM runtime.</td>
    <td><code>2</code></td>
    <td><span style="color:#ff8888">If it does not exceed 2GB, the server may not be running.</span></td>
  </tr>
  <tr>
    <td><code>enable_traffic_monitor</code></td>
    <td>Sets whether the server traffic measurement feature is enabled.</td>
    <td><code>false</code></td>
    <td><span style="color:#ff8888">The feature is still an alpha state. See the README.MD file for a list of known issues.</span></td>
  </tr>
  <tr>
    <td><code>traffic_redirect_port</code></td>
    <td>You must redirect the port to use the Server Traffic Measurement feature. Set the port to use when redirecting.</td>
    <td><code>25555</code></td>
    <td><span style="color:#ff8888">The feature is still an alpha state. See the README.MD file for a list of known issues. If the <code>enable_traffic_monitor</code> option is turned off, this option is ignored.</span></td>
  </tr>
  <tr>
    <td><code>api_port</code></td>
    <td>Sets the port that is exposed to the outside.</td>
    <td><code>8080</code></td>
    <td><span>N/A</span></td>
  </tr>
  <tr>
    <td><code>plugin_api_port</code></td>
    <td>Please make it the same as the port in the Dashify plugin setting that goes into the server. Dashify retrieves information through that port. If the ports do not match, proper information is not delivered.</td>
    <td><code>8081</code></td>
    <td><span>The plug-in settings file is located in the path 'plugins/Dashify/config.yml'.</span></td>
  </tr>
  <tr>
    <td><code>plugins</code></td>
    <td>Please write down the download link of the plug-in to enter the server in array form.</td>
    <td><code>["https://github.com/MC-Dashify/plugin/releases/latest/download/dashify-plugin-all.jar"]</code></td>
    <td><span>If the link is incorrect, it is ignored. The plug-ins listed here are downloaded to the latest version each time launcher is started.</span></td>
  </tr>
  <tr>
    <td><code>jar_args</code></td>
    <td>Set the arguments to run the server.</td>
    <td><code>["nogui"]</code></td>
    <td><span style="color:#ff8888">Invalid arguments are ignored. If you enter some options incorrectly, the server may not function properly.</span></td>
  </tr>
</tbody>
</table>
