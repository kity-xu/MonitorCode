package protocol

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/rgamba/evtwebsocket"
	"haina.im/monitor/monitor_node/share"
	//"time"
)

type SocketClient struct {
	C evtwebsocket.Conn

	Received []byte
}

func (this *SocketClient) Client(ip, port string) {
	url := "ws://" + ip + ":" + port + "/socket"
	this.C = evtwebsocket.Conn{

		// When connection is established
		OnConnected: func(w *evtwebsocket.Conn) {
			l4g.Debug("Connected")
		},

		// When a message arrives
		OnMessage: func(msg []byte, w *evtwebsocket.Conn) {
			//l4g.Debug("OnMessage: %s\n", msg)
			this.Received = msg
		},

		// When the client disconnects for any reason
		OnError: func(err error) {
			l4g.Debug("** ERROR **\n%s\n", err.Error())
		},

		// This is used to match the request and response messagesP>termina
		MatchMsg: func(req, resp []byte) bool {
			return string(req) == string(resp)
		},

		// Auto reconnect on error
		Reconnect: true,

		// Set the ping interval (optional)
		PingIntervalSecs: 5,

		// Set the ping message (optional)
		PingMsg: []byte("PING"),
	}

	// Connect
	if err := this.C.Dial(url, ""); err != nil {
		l4g.Error(err)
	}
}

func (this *SocketClient) Send(ss []byte) {
	if !this.C.IsConnected() {
		l4g.Error("错误码：%d", share.WEBSOCKET_DISCONNECTED)
		return
	}
	msg := evtwebsocket.Msg{
		Body: ss,
		Callback: func(resp []byte, w *evtwebsocket.Conn) {
			l4g.Debug("Callback: %s\n", resp)
		},
	}

	if err := this.C.Send(msg); err != nil {
		l4g.Debug("Unable to send: ", err.Error())
	}
	//time.Sleep(2000)
}

func (this *SocketClient) Close() {

}
