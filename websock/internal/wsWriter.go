package internal

import (
	"fmt"
	"sync"

	"github.com/SarthakJha/distr-websock/models"
	"github.com/SarthakJha/distr-websock/repository"
	"github.com/gorilla/websocket"
)

func WSWriterHandler(writeBuff chan models.Message, sockMap *sync.Map, db repository.MessageRepository) {
	for {
		select {
		case msg := <-writeBuff:
			// query local sync.MAP for key:id to get ws.Sock
			sock, _ := sockMap.Load(msg.RecieverID)
			if sock == nil {
				sockMap.Delete(msg.RecieverID)
				continue
			}
			websock, ok := sock.(*websocket.Conn) // converting interface to a struct
			if !ok {
				continue
			}

			// write to the return ws.Conn
			err := websock.WriteJSON(msg)
			if err != nil {
				fmt.Println(err.Error())
				sockMap.Delete(msg.RecieverID)
			}

			// set status to 'DELIVERED'
			err = db.SetStatusToDelivered(msg.MessageID)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
