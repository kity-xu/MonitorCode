package protocol

import (
	"bytes"
	"encoding/gob"
	l4g "github.com/alecthomas/log4go"
	"github.com/rgamba/evtwebsocket"
	"time"
)

func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

func WebsocketClient(cc chan ResultData) {
	//var origin = "http://xiaodong.xiaodong.im/"
	var url = "ws://192.168.2.79:5010/socket"
	c := evtwebsocket.Conn{

		// When connection is established
		OnConnected: func(w *evtwebsocket.Conn) {
			l4g.Debug("Connected")
		},

		// When a message arrives
		OnMessage: func(msg []byte, w *evtwebsocket.Conn) {
			l4g.Debug("OnMessage: %s\n", msg)
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
	if err := c.Dial(url, ""); err != nil {
		l4g.Error(err)
	}

	// Create the message with a callback
	//context := "[\"pusher:subscribe\", {\"auth\":\"\", \"channel\":\"channel1\"}]"
	context, err := Encode(<-cc)
	if err != nil {
		l4g.Error("序列化失败:%s", err)
	}
	msg := evtwebsocket.Msg{
		Body: context,
		Callback: func(resp []byte, w *evtwebsocket.Conn) {
			l4g.Debug("Callback: %s\n", resp)
		},
	}

	l4g.Debug("Sending message: %s\n", msg.Body)

	// Send the message to the server
	if err := c.Send(msg); err != nil {
		l4g.Debug("Unable to send: ", err.Error())
	}

	// Take a break
	time.Sleep(time.Second * 3)
}
