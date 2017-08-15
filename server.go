package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"encoding/json"
)

var channels [10]chan Event

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func receive(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	payload := r.Form["payload"][0]
	log.Println(payload)

	var event Event
	if err := json.Unmarshal([]byte(payload), &event); err != nil {
		log.Fatal("unmarshal:", err)
	}

	for _, channel := range channels {
		if channel != nil {
			channel <- event
		}
	}

	fmt.Fprintf(w, "OK")
}

func register() chan Event {
	for i, v := range channels {
		if v == nil { 
			channels[i] = make(chan Event)
			return channels[i]
		}
	}
	return nil
}

func unregister(channel chan Event) {
	for i, v := range channels {
		if v == channel {
			channels[i] = nil
		}
	}
}

func socket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	var channel = register()
	if channel == nil {
		log.Println("register: no more connections")
		conn.Close()
		return
	}
	go func() {
		defer conn.Close()
		for {
			event := <- channel
			if err := conn.WriteJSON(event); err != nil {
				log.Println("write:", err)
				unregister(channel)
				break
			}
		}
	}()
}

func Serve() {
	http.HandleFunc("/slack", receive)
	http.HandleFunc("/ws", socket)
	http.ListenAndServe(":80", nil)
}
