package main

import (
	"fmt"
	"haina.im/monitor/monitor_centre/protocol"
	"time"
)

func main() {
	server := new(protocol.SocketServer)
	go server.StartSocketserver()

	for {
		//if server.Buffer.Time != "" { //判断结构体是否为空，以防止取值越界
		//	fmt.Println("-------------", server.Buffer.Apps[0].Name)
		//}
		fmt.Println("-------------", server.Buffer)
		time.Sleep(time.Duration(5) * time.Second)
	}
	//message := "Good luck..."
	//protocol.SendMail(message)
}
