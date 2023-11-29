package main

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

var natsConnection *nats.Conn //, _ = nats.Connect("nats://127.0.0.1:14222")

func Subscribe() {

}

func HandleResponse(data any) {
	natsConnection, _ = nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	log.Println("Sending Message...")
	// Publish message on subject
	// message := json.Unmarshal([]byte(data), string)
	message, _ := json.Marshal(data)
	natsConnection.Publish("TPOUT", message)
}
