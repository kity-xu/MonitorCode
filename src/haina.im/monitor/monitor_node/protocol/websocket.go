package protocol

import (
	"bytes"
	"encoding/gob"
	l4g "github.com/alecthomas/log4go"
	"github.com/rgamba/evtwebsocket"
	"haina.im/monitor/monitor_node/share"
	"time"
)

type Applications struct {
	As []ResultApp
}

type Data struct {
	Apps interface{}
	Exp  interface{}
	Osys interface{}
	Time string
}

type Res struct {
	Apps interface{}
	Exp  interface{}
	Osys interface{}
	Time string
}

func Encode(da MonitorData) ([]byte, error) {
	var buf bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&buf) // Will write to network.

	gob.Register(Applications{})
	gob.Register(SysStatus{})
	gob.Register(Osystem{})

	apps := Applications{da.Apps}
	last := Data{apps, da.Statu, da.Osys, da.Time}

	err := enc.Encode(last)
	if err != nil {
		return nil, err
	}

	// dec := gob.NewDecoder(&buf)
	// var res Res
	// e := dec.Decode(&res)
	// if e != nil {
	// 	l4g.Debug("序列化失败:%s", err)
	// }
	// l4g.Debug("Decode-----", res)

	return buf.Bytes(), nil
}

// func Encode(data interface{}) ([]byte, error) {
// 	buf := bytes.NewBuffer(nil)
// 	enc := gob.NewEncoder(buf)
// 	err := enc.Encode(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }

func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)

	// gob.Register(Applications{})
	// gob.Register(Explorer{})
	// gob.Register(Osystem{})

	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

type SocketClient struct {
	C evtwebsocket.Conn

	Received []byte
}

func (this *SocketClient) WebsocketClient() {
	//var origin = "http://xiaodong.xiaodong.im/"
	var url = "ws://192.168.2.79:5010/socket"
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

	// context, err := Encode(cc)
	// if err != nil {
	// 	l4g.Error("序列化失败:%s", err)
	// }

	// l4g.Debug("---websocket中最终的数据：", context)

	// var R Res
	// // buf := bytes.NewBuffer(context)
	// // dec := gob.NewDecoder(buf)

	// e2 := Decode(context, &R)
	// if e2 != nil {
	// 	l4g.Error("序列化失败:%s", e2)
	// }
	// l4g.Debug("---Result:", R)

	//l4g.Debug("----c----------:%t", this.C)

	//l4g.Debug("Sending message: %v\n", msg)

	// Send the message to the server

	// Take a break
	//time.Sleep(time.Second * 3)
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
	time.Sleep(2000)
}
