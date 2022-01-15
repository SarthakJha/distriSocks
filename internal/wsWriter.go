package internal

import (
	"github.com/SarthakJha/distr-websock/models"
)

func WSWriterHandler(writeBuff chan models.Message) {
	for {
		select {
		case msg := <-writeBuff:
			// query local sync.MAP for key:id
			// write to the return ws.Conn
			// set status to 'DELIVERED'
		}
	}
}
