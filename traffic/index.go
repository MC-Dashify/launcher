package traffic

import (
	"fmt"
	"io"
	"net"

	"github.com/MC-Dashify/launcher/config"
	"github.com/MC-Dashify/launcher/global"
	"github.com/MC-Dashify/launcher/utils/logger"
)

func StartTrafficMonitor() {
	listenAddress := fmt.Sprintf("0.0.0.0:%d", config.ConfigContent.TrafficRedirectPort) // 리다이렉트할 주소
	redirectAddress := fmt.Sprintf("localhost:%d", global.MCOriginPort)                  // 리다이렉트할 대상 주소

	// TCP 리스너 생성
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		logger.Fatal(fmt.Sprintf("[TrafficMonitor] Error listening on %s: %v", listenAddress, err))
	}

	logger.Info(fmt.Sprintf("[TrafficMonitor] Redirecting traffic from %s to %s", listenAddress, redirectAddress))

	// go func() {
	// 	for range time.Tick(time.Second) {
	// 		global.TrafficClientsMutex.RLock()
	// 		for clientAddr, stats := range global.TrafficClients {
	// 			receivedMbps := float64(stats.ReceivedBytes*8) / (1000 * 1)
	// 			sentMbps := float64(stats.SentBytes*8) / (1000 * 1)
	// 			logger.Debug(fmt.Sprintf("[TrafficMonitor] Client %s\t\tReceived Traffic: %.2f Kbps\tSent Traffic: %.2f Kbps\n", clientAddr, receivedMbps, sentMbps))
	// 			stats.ReceivedBytes = 0
	// 			stats.SentBytes = 0
	// 		}
	// 		global.TrafficClientsMutex.RUnlock()
	// 	}
	// }()

	for {
		// 클라이언트 연결 대기
		clientConn, err := listener.Accept()
		if err != nil {
			logger.Warn(fmt.Sprintf("[TrafficMonitor] Error accepting connection: %v", err))
			continue
		}

		// 대상 서버에 연결
		redirectConn, err := net.Dial("tcp", redirectAddress)
		if err != nil {
			logger.Warn(fmt.Sprintf("[TrafficMonitor] Error connecting to redirect address: %v", err))
			clientConn.Close()
			continue
		}

		// 클라이언트 주소 추출
		clientAddr := clientConn.RemoteAddr().String()

		// 트래픽 계산을 위한 변수
		stats := &global.TrafficClientStats{}

		// 클라이언트에서 대상 서버로 데이터 복사
		go func(clientConn net.Conn, redirectConn net.Conn, clientAddr string, stats *global.TrafficClientStats) {
			buffer := make([]byte, 8192) // 복사 버퍼
			for {
				n, err := clientConn.Read(buffer)
				if err != nil {
					if err != io.EOF {
						logger.Info(fmt.Sprintf("[TrafficMonitor] Client %s connection closed!", clientAddr))
					}
					break
				}

				_, err = redirectConn.Write(buffer[:n])
				if err != nil {
					logger.Warn(fmt.Sprintf("[TrafficMonitor] Error writing to redirect: %v", err))
					break
				}

				stats.ReceivedBytes += int64(n)
			}

			clientConn.Close()
			redirectConn.Close()

			global.TrafficClientsMutex.Lock()
			delete(global.TrafficClients, clientAddr)
			global.TrafficClientsMutex.Unlock()
		}(clientConn, redirectConn, clientAddr, stats)

		// 대상 서버에서 클라이언트로 데이터 복사
		go func(clientConn net.Conn, redirectConn net.Conn, clientAddr string, stats *global.TrafficClientStats) {
			buffer := make([]byte, 8192) // 복사 버퍼
			for {
				n, err := redirectConn.Read(buffer)
				if err != nil {
					if err != io.EOF {
						logger.Info(fmt.Sprintf("[TrafficMonitor] Client %s connection closed!", clientAddr))
					}
					break
				}

				_, err = clientConn.Write(buffer[:n])
				if err != nil {
					logger.Warn(fmt.Sprintf("[TrafficMonitor] Error writing to client: %v", err))
					break
				}

				stats.SentBytes += int64(n)
			}

			clientConn.Close()
			redirectConn.Close()
		}(clientConn, redirectConn, clientAddr, stats)

		// 클라이언트 정보 저장
		global.TrafficClientsMutex.Lock()
		global.TrafficClients[clientAddr] = stats
		global.TrafficClientsMutex.Unlock()
	}
}
