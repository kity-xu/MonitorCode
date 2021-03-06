package main

import (
	"fmt"
	"github.com/cihub/seelog"
	"haina.im/monitor/monitor_centre/protocol"
	"time"
)

func main() {
	//notif := &protocol.Notification{}
	log, _ := seelog.LoggerFromConfigAsFile("haina.im/monitor/monitor_centre/config/log.xml")
	defer log.Flush()

	server := new(protocol.SocketServer)
	go server.StartSocketserver()
	//var armmsg string

	for {
		//fmt.Println("---------MonitorNode:%v", server.Buffer)
		// if server.IsConn && server.Buffer.Time != "" { //判断结构体是否为空，以防止取值越界
		// 	if len(server.Buffer.Arms) > 0 {
		// 		for _, arm := range server.Buffer.Arms {
		// 			armmsg += fmt.Sprintf("错误码：%d; 错误级别：%d; 错误信息：%s --时间：%s\n", arm.Code, arm.Grade, arm.Msg, server.Buffer.Time)
		// 		}
		// 		log.Error(armmsg) //日志落地
		// 		//notif.SendMail(armmsg) //发邮件
		// 		armmsg = ""
		// 		break
		// 	}
		// }
		if len(server.Conn) > 0 {
			for k, v := range server.Conn {
				fmt.Println("key:", k)
				fmt.Println("value:", v.Data)
			}
		}

		time.Sleep(time.Duration(5) * time.Second)
	}

	//message := "Good luck..."
	//protocol.SendMail(armmsg)
}
