package internal

import (
	"log"
	"net/http"

	"github.com/SarthakJha/distr-websock/models"
	"github.com/gorilla/websocket"
)

type Chans struct {
	publish *chan models.Message
}

func (c *Chans) InitChan(buffer int64) {
	if buffer < 1 {
		log.Fatalln("buffr size should be greater than one")
	}
	ch := make(chan models.Message, buffer)
	c.publish = &ch
}

func (c *Chans) HandleConn(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	// save to map [username:ws.Conn]
	// save to redis [user_id:pod_id]
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer ws.Close()

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			// TODO:- handle error

			// on disconnect:
			// 1. remove key from sync.Map
			// 2. remove key from redis
		}
		if c.publish == nil {
			log.Fatalln("channel not initialised")
		}
		*(c.publish) <- msg
	}
}
