package protocol

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

type MonitorData struct {
	Apps   []AppStatus
	Num    int
	Status SysStatus
	Info   SysInfo
	Time   string
}

type SocketServer struct {
	Buffer MonitorData
}

func (this *SocketServer) Echo(ws *websocket.Conn) {
	var err error
	this.Buffer = MonitorData{}
	for {

		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		//fmt.Println("----------bytes------------:\n", reply)

		err := json.Unmarshal([]byte(reply), &this.Buffer)

		//fmt.Println("----------bengin------------:\n", this.Buffer)

		msg := "googd luck xulang"
		if err = websocket.Message.Send(ws, msg); err != nil {
			break
		}
	}
}

func (this *SocketServer) StartSocketserver() {
	fmt.Println("begin")
	http.Handle("/", http.FileServer(http.Dir("."))) // <-- note this line

	http.Handle("/socket", websocket.Handler(this.Echo))

	if err := http.ListenAndServe(":5010", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	fmt.Println("end")
}
