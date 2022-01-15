package models

type Payload struct {
	Msg        string `json:"msg"`
	SenderID   string `json:"sender_id"`
	RecieverID string `json:"reciever_id"`
}
