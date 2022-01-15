package models

import "github.com/gorilla/websocket"

type WSWriterBuffer struct {
	Payload
	WSConnection websocket.Conn
}

type Payload struct {
	Msg        string `json:"msg"`
	SenderID   string `json:"sender_id"`
	RecieverID string `json:"reciever_id"`
}
