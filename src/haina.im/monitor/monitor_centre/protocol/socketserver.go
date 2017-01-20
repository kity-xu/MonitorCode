package protocol

import (
	//"container/list"
	"encoding/json"
	"fmt"
	"github.com/cihub/seelog"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/websocket"
)

type Record struct {
	Name string
	Data string
}

type AppStatus struct {
	Name    string
	Num     int
	Records []Record
}

type SysStatus struct {
	Tasks string
	Cpu   string
	Mem   string
	Swap  string
}

type SysInfo struct {
	IP   string
	ID   string
	Name string
}

type Alarm struct {
	IsErr bool
	Code  int
	Msg   string
	Grade int
}

type MonitorData struct {
	Info   SysInfo
	Status SysStatus
	Num    int
	Apps   []AppStatus
	Arms   []Alarm
	Time   string
}

type SocketServer struct {
	Buffer MonitorData
	IsConn bool
}

type Client struct {
	ws       *websocket.Conn
	Addr     string
	Data     MonitorData
	LastData MonitorData
}

var logf seelog.LoggerInterface

var clients map[string]Client

func (this *SocketServer) websocketServer(ws *websocket.Conn) {
	var (
		err       error
		clientMsg string
	)

	id := ws.Request().RemoteAddr
	ip := strings.Split(id, ":")[0]
	node := Client{}
	node.Addr = id
	node.ws = ws

	defer func() {
		if err = ws.Close(); err != nil {
			logf.Error("Websocket could not be closed", err.Error())
		}

		logf.Flush()
		logf.Close()
	}()

	for {
		if err = websocket.Message.Receive(ws, &clientMsg); err != nil {
			logf.Debug(id, "receive disconnected...")
			return
		}

		node.LastData = node.Data
		if err := json.Unmarshal([]byte(clientMsg), &node.Data); err != nil {
			logf.Error(id, " 反序列化错误...")
			websocket.Message.Send(ws, "反序列化错误。。。")
			return
		}

		fmt.Println("=----------addrIP:", id)
		fmt.Println("=----------addrIP:", ip)

		clients[ip] = node

		msg := "good luck xulang"
		if err = websocket.Message.Send(ws, msg); err != nil {
			logf.Debug(id, "send disconnected...")
		}
		logf.Debug(clients)
		fmt.Println(len(clients))
	}
}

func (this *SocketServer) StartSocketserver() {
	clients = make(map[string]Client)
	logf, _ = seelog.LoggerFromConfigAsFile("haina.im/monitor/monitor_centre/config/logconfig.xml")

	fmt.Println("begin")
	http.Handle("/", http.FileServer(http.Dir("."))) // <-- note this line

	http.Handle("/socket", websocket.Handler(this.websocketServer))

	if err := http.ListenAndServe(":5010", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	fmt.Println("end")
}
