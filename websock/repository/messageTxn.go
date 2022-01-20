package repository

import (
	"time"

	"github.com/SarthakJha/distr-websock/internal/constants"
	"github.com/SarthakJha/distr-websock/internal/models"
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
	message.Status = constants.MSG_STATUS_NONE
	err := msg.Table.Put(message).Run()
	return err
}

func (msg *MessageRepository) SetStatusToSent(messageID string) error {
	err := msg.Table.Update("message_id", messageID).Set("status", constants.MSG_STATUS_SENT).Run()
	if err != nil {
		return err
	}
	return nil
}

func (msg *MessageRepository) SetStatusToDelivered(messageID string) error {
	err := msg.Table.Update("message_id", messageID).Set("status", constants.MSG_STATUS_DELIVERED).Run()
	if err != nil {
		return err
	}
	return nil
}
