package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/SarthakJha/distr-websock/internal/models"
	"github.com/SarthakJha/distr-websock/repository"
	"github.com/gorilla/websocket"
)

type Chans struct {
	publish *chan models.Message
	sockMap *sync.Map
	db      *repository.ConnectionRepository
}

// chan getter
func (c *Chans) GetPubChan() *chan models.Message {
	return c.publish
}
func (c *Chans) GetMap() *sync.Map {
	return c.sockMap
}

func (c *Chans) InitChan(buffer int64) {
	if buffer < 1 {
		log.Fatalln("buffr size should be greater than one")
	}
	ch := make(chan models.Message, buffer)
	var m sync.Map
	c.sockMap = &m
	c.publish = &ch
}

func (c *Chans) HandleConn(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// save to map [username:ws.Conn]
	c.sockMap.Store("user-id from cookie", ws)

	// save to redis [user_id:pod_id]
	c.db.SetWSConnection("user-id from cookie", os.Getenv("POD_ID"))

	defer ws.Close()
	defer close(*c.publish)

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			// TODO:- handle error

			// on disconnect:
			// 1. remove key from sync.Map
			c.sockMap.Delete("user-id from cookie")
			// 2. remove key from redis
			err1 := c.db.DeleteWSConnection("user-id from cookie")
			if err1 != nil {
				fmt.Println(err1.Error())
			}

		}
		if c.publish == nil {
			log.Fatalln("channel not initialised")
		}
		*(c.publish) <- msg
	}
}
