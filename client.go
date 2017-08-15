package main

import (
	"log"
	"github.com/gorilla/websocket"
	"net/http"
)

func Connect(host string) {

	var dialer = websocket.Dialer {
		Proxy: http.ProxyFromEnvironment,
	}
	
	conn, _, err := dialer.Dial("ws://" + host + "/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()
	var event Event

	for {
		if err := conn.ReadJSON(&event); err != nil {
			log.Println("read:", err)
		}
		log.Printf("recv: %s", event)
	}

}