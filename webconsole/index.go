package webconsole

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/MC-Dashify/launcher/config"
	"github.com/MC-Dashify/launcher/global"
	"github.com/MC-Dashify/launcher/i18n"
	"github.com/MC-Dashify/launcher/utils"
	"github.com/MC-Dashify/launcher/utils/logger"
	"github.com/gorilla/websocket"
)

var (
	Server *http.Server
	h      *WebSocketHandler // WebSocketHandler 인스턴스를 전역 변수로 변경
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var IsRestart bool = false

type WebSocketHandler struct {
	cmdIn        *io.WriteCloser
	cmdOut       *io.ReadCloser
	cmdErr       *io.ReadCloser
	console      *log.Logger
	inputChan    chan string
	outputChan   chan string
	connections  []*websocket.Conn
	mu           sync.Mutex
	commandEnded chan bool // command 종료 신호를 전달하는 채널
	connectChan  chan *websocket.Conn
	disconnectCh chan *websocket.Conn
}

func (h *WebSocketHandler) ReadCommandOutput() {
	scanner := bufio.NewScanner(*h.cmdOut)
	for scanner.Scan() {
		output := scanner.Text()
		h.mu.Lock()
		// if len(h.connections) > 0 {
		// 	logger.Debug(fmt.Sprintf("%v", h.connections)) // 커넥션 주소. 동일한 주소 여러개인경우 메세지 여러번 나감
		// }
		for _, conn := range h.connections {
			err := conn.WriteMessage(websocket.TextMessage, []byte(output))
			if err != nil {
				logger.Error(fmt.Sprintf("%+v", err))
			}
		}
		h.mu.Unlock()
		h.console.Println(string(output)) // 콘솔에 출력
	}
	if err := scanner.Err(); err != nil {
		logger.Error(fmt.Sprintf("%+v", err))
	}
}

func (h *WebSocketHandler) ReadCommandError() {
	scanner := bufio.NewScanner(*h.cmdErr)
	for scanner.Scan() {
		output := scanner.Text()
		h.mu.Lock()
		for _, conn := range h.connections {
			err := conn.WriteMessage(websocket.TextMessage, []byte(output))
			if err != nil {
				logger.Error(fmt.Sprintf("%+v", err))
			}
		}
		h.mu.Unlock()
		h.console.Println(output) // 콘솔에 출력
		// logger.Debug(fmt.Sprintf("%+v", output))
	}
	if err := scanner.Err(); err != nil {
		logger.Error(fmt.Sprintf("%+v", err))
	}
}

func (h *WebSocketHandler) ReadCommands() {
	for input := range h.inputChan {
		_, err := (*h.cmdIn).Write([]byte(input + "\n"))
		if err != nil {
			logger.Error(fmt.Sprintf("%+v", err))
		}
	}
}

func (h *WebSocketHandler) WriteCommand(command string) {
	// logger.Debug(fmt.Sprintf("%v", h.cmdIn))
	_, err := (*h.cmdIn).Write([]byte(command))
	if err != nil {
		logger.Error(fmt.Sprintf("%+v", err))
	}
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// CheckOrigin 함수를 통해 모든 원천(Origin)에서의 웹소켓 연결을 허용합니다.
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("%+v", err))
		return
	}
	logger.Info(strings.ReplaceAll(i18n.Get("webconsole.connection.opened"), "$connection", conn.RemoteAddr().String()))

	// 새로운 웹소켓 연결을 connectChan에 전달
	// logger.Debug(fmt.Sprintf("%v", &h.connectChan))
	h.connectChan <- conn

	defer func() {
		// 웹소켓 연결이 종료될 때 disconnectCh에 전달
		h.disconnectCh <- conn
		h.mu.Lock()
		for i, c := range h.connections {
			if c == conn {
				h.connections = append(h.connections[:i], h.connections[i+1:]...)
				break
			}
		}
		h.mu.Unlock()

		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseMessage, websocket.CloseNoStatusReceived) {
				logger.Error(strings.ReplaceAll(strings.ReplaceAll(i18n.Get("webconsole.connection.closed.error"), "$connection", conn.RemoteAddr().String()), "$error", err.Error()))
			}
			break
		}
		// 메시지를 한 번만 처리하기 위해 다음과 같이 수정합니다.
		if !utils.Contains(config.ConfigContent.WebConsoleDisabledCmds, string(msg)) {
			h.inputChan <- string(msg)
			logger.Info(strings.ReplaceAll(strings.ReplaceAll(i18n.Get("webconsole.connection.cmd.received"), "$remote", conn.RemoteAddr().String()), "$command", string(msg)))
		}
	}
}

func StartWebsocket() {
	logger.Info(i18n.Get("webconsole.started1"))
	logger.Info(i18n.Get("webconsole.started2"))
	logger.Info(i18n.Get("webconsole.started1"))
	logger.Debug(i18n.Get("webconsole.chk.valid.prev.connection"))
	if h != nil {
		logger.Debug(strings.ReplaceAll(i18n.Get("webconsole.restoring.prev.connection"), "$connection", fmt.Sprintf("%+v", &h.connectChan)))
	}
	h = &WebSocketHandler{
		console:      log.New(os.Stdout, "", 0),
		inputChan:    make(chan string),
		outputChan:   make(chan string),
		connections:  make([]*websocket.Conn, 0),
		commandEnded: make(chan bool),
		connectChan:  make(chan *websocket.Conn),
		disconnectCh: make(chan *websocket.Conn),
	}
	// logger.Debug(fmt.Sprintf("%+v", &h.connectChan))
	go func() {
		for conn := range h.connectChan {
			h.mu.Lock()
			h.connections = append(h.connections, conn)
			h.mu.Unlock()
		}
	}()

	go func() {
		for conn := range h.disconnectCh {
			h.mu.Lock()
			for i, c := range h.connections {
				if c == conn {
					h.connections = append(h.connections[:i], h.connections[i+1:]...)
					break
				}
			}
			h.mu.Unlock()
			conn.Close()
			logger.Info(strings.ReplaceAll(i18n.Get("webconsole.connection.closed"), "$connection", conn.RemoteAddr().String()))
		}
	}()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				if global.IsMCServerRunning {
					logger.Warn(i18n.Get("general.unsafe.shutdown"))
				}
			}
			input = strings.TrimSpace(input)
			h.inputChan <- input
		}
	}()

	cmdIn, err := global.Cmd.StdinPipe()
	if err != nil {
		logger.Error(fmt.Sprintf("%+v", err))
		log.Fatal(err)
	}

	cmdOut, err := global.Cmd.StdoutPipe()
	if err != nil {
		logger.Error(fmt.Sprintf("%+v", err))
		log.Fatal(err)
	}

	cmdErr, err := global.Cmd.StderrPipe()
	if err != nil {
		logger.Error(fmt.Sprintf("%+v", err))
		log.Fatal(err)
	}
	h.cmdIn = &cmdIn
	h.cmdOut = &cmdOut
	h.cmdErr = &cmdErr

	go h.ReadCommandOutput()
	go h.ReadCommandError()
	go h.ReadCommands()

	err = global.Cmd.Start()
	if err != nil {
		logger.Fatal(strings.ReplaceAll(i18n.Get("java.jvm.start.failed"), "$error", err.Error()))
	} else {
		global.IsMCServerRunning = true
		global.NormalStatusExit = true
	}

	go func() {
		err := global.Cmd.Wait()
		if err != nil {
			global.IsMCServerRunning = false
			global.NormalStatusExit = false
			logger.Debug(fmt.Sprintf("%+v", err))
		} else {
			global.IsMCServerRunning = false
			global.NormalStatusExit = true
		}
		logger.Debug(i18n.Get("java.jvm.stopped"))
		h.commandEnded <- true // command 종료 신호를 전달
	}()

	go func() {
		for range h.commandEnded {
			h.mu.Lock()
			for _, conn := range h.connections {
				err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					logger.Warn(strings.ReplaceAll(strings.ReplaceAll(i18n.Get("webconsole.connection.close.msg.send.fail"), "$connection", conn.RemoteAddr().String()), "$error", err.Error()))
				}
				conn.Close()
			}
			h.mu.Unlock()

			// Server.Shutdown(context.Background())
			close(h.inputChan)
			close(h.outputChan)
			close(h.commandEnded)
			close(h.connectChan)
			close(h.disconnectCh)
		}
	}()

	if !IsRestart {
		http.HandleFunc("/console", HandleWebSocket)
		http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("Server is alive."))
		})
	}
}

func RunServerWithWebSocketHandler(arguments []string, connections []*websocket.Conn) {
	MCServerLauncher("java", arguments)

	if connections != nil {
		h.connections = connections // 이전 연결 복원
	}
	StartWebsocket()
}

func RunServer(arguments []string) {
	RunServerWithWebSocketHandler(global.JarArgs, nil)
}

func MCServerLauncher(baseCmd string, cmdArgs []string) error {
	if global.Cmd != nil && global.Cmd.Process != nil {
		_ = global.Cmd.Process.Kill()
	}

	logger.Debug(strings.ReplaceAll(i18n.Get("general.exec"), "$command", baseCmd+" "+strings.Join(cmdArgs, " ")))

	global.Cmd = exec.Command(baseCmd, cmdArgs...)
	global.Cmd.Env = os.Environ()

	return nil
}
