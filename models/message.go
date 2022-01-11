package models

import "time"

type Message struct {
	Payload    string    `dynamo:"payload"`
	SenderID   string    `dynamo:"sender_id"`
	RecieverID string    `dynamo:"recv_id"`
	CreatedAt  time.Time `dynamo:"created_at"`
	MessageID  string    `dynamo:"message_id"`
}
