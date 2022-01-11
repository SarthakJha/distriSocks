package repository

import (
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
	err := msg.Table.Put(message).Run()
	return err
}
