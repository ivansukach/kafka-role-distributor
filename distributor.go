package main

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"time"
)

func main() {
	selector := false
	origin := "http://localhost/"
	urlS1 := "ws://localhost:1333/ws"
	urlS2 := "ws://localhost:1323/ws"
	wss1, err := websocket.Dial(urlS1, "", origin)
	if err != nil {
		log.Error(err)
	}
	wss2, err := websocket.Dial(urlS2, "", origin)
	if err != nil {
		log.Error(err)
	}
	tRole := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-tRole.C:
			if selector {
				if _, err := wss1.Write([]byte("server2:producer")); err != nil {
					log.Fatal(err)
				}
				if _, err := wss2.Write([]byte("server2:producer")); err != nil {
					log.Fatal(err)
				}
				log.Println("server2:producer")
				selector = false
			} else {
				if _, err := wss1.Write([]byte("server1:producer")); err != nil {
					log.Fatal(err)
				}
				if _, err := wss2.Write([]byte("server1:producer")); err != nil {
					log.Fatal(err)
				}
				log.Println("server1:producer")
				selector = true
			}
		}
	}
}
