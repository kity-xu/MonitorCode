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
		//fmt.Println("---------MonitorNode:%v", server.Buffer)
		if server.Buffer.Time != "" { //判断结构体是否为空，以防止取值越界
			if len(server.Buffer.Arms) > 0 {
				for _, arm := range server.Buffer.Arms {
					armmsg := fmt.Sprintf("错误码：%d; 错误级别：%d; 错误信息：%s --时间：%s\n", arm.Code, arm.Grade, arm.Msg, server.Buffer.Time)
					protocol.SendMail(armmsg)
					break
				}
			}
		}
		time.Sleep(time.Duration(5) * time.Second)
	}

	//message := "Good luck..."
	//protocol.SendMail(armmsg)
}
