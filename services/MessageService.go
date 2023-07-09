package services

import (
	"github.com/gorilla/websocket"
	"log"
)

type MessageService struct {
}

type MessageJson struct {
	Code    int           `json:"code"`
	Data    []interface{} `json:"data"`
	Message string        `json:"message"`
}

func (Me *MessageService) LogAdd(ws *websocket.Conn, message *RequestJson) {
	// struct to byte
	log.Printf("%#v", message.Data)
	for _, v := range message.Data {
		log.Printf("%#v", v)
	}
	err := ws.WriteMessage(1, []byte("我是LogAdd"))
	if err != nil {
		return
	}
}

func (Me *MessageService) ISun(ws *websocket.Conn, message *RequestJson) {
	// struct to byte
	err := ws.WriteMessage(1, []byte("我是ISun"))
	if err != nil {
		return
	}
}

func (Me *MessageService) SendAct(ws *websocket.Conn, message *RequestJson) {
	// struct to byte
	err := ws.WriteMessage(1, []byte("我是LogList"))
	if err != nil {
		return
	}
}
