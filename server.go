package main

import (
	"fmt"
	"log"
	"net/http"
	"bytes"
	"encoding/gob"

	"golang.org/x/net/websocket"
)

type ResultRec struct {
	Name string
	Data string
}

type ResultApp struct {
	Name string
	Recs []ResultRec
}

type Explorer struct {
	Tasks string
	Cpu   string
	Mem   string
	Swap  string
}

type Osystem struct {
	Sys      string
	User     string
	IP       string
	Version  string
	Platform string
}

type ResultData struct {
	Apps []ResultApp
	Exp  Explorer
	Osys Osystem
}

func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

func Echo(ws *websocket.Conn) {
	var err error

	for {

		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		var to ResultData
		Decode([]byte(reply), to)

		fmt.Printf("Received back from client: %s" + to.Exp.Cpu)
		fmt.Println("Received len:", len(reply))

		msg := "Received = " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	fmt.Println("begin")
	http.Handle("/", http.FileServer(http.Dir("."))) // <-- note this line

	http.Handle("/socket", websocket.Handler(Echo))

	if err := http.ListenAndServe(":5010", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	fmt.Println("end")
}
