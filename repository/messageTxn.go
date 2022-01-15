package repository

import (
	"time"

	"github.com/SarthakJha/distr-websock/models"
	"github.com/gofrs/uuid"
)

func (msg *MessageRepository) SaveMessage(message models.Message) error {
	if message.MessageID == "" {
		id, err := uuid.NewV1()
		if err != nil {
			return err
		}
		message.MessageID = id.String()
	}
	message.CreatedAt = time.Now()
	err := msg.Table.Put(message).Run()
	return err
}
