package models

import "time"

type Message struct {
	Payload    string    `dynamo:"payload" json:"payload"`
	SenderID   string    `dynamo:"sender_id" json:"sender_id"`
	RecieverID string    `dynamo:"recv_id" json:"reciever_id"`
	CreatedAt  time.Time `dynamo:"created_at,unixtime" json:"created_at,omitempty" `
	MessageID  string    `dynamo:"message_id" json:"message_at,omitempty"`
	Status     string    `dynamo:"status" json:"status,omitempty"`
}
