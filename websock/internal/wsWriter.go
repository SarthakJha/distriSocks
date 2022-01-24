package internal

import (
	"fmt"
	"sync"

	"github.com/SarthakJha/distr-websock/internal/models"
	"github.com/SarthakJha/distr-websock/repository"
	"github.com/gorilla/websocket"
)

// consumer side of the channel
func WSWriterHandler(writeBuff chan models.Message, sockMap *sync.Map, db repository.MessageRepository, wg *sync.WaitGroup) {

	defer wg.Done() // gracefull shutdown of consumer worker
	// this will automatically shutdown when producing side of channel shuts
	for msg := range writeBuff {
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
