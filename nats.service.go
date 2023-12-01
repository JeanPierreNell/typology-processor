package main

import (
	"log"

	P "typology-processor/proto"

	"google.golang.org/protobuf/proto"

	"github.com/nats-io/nats.go"
)

var natsConnection *nats.Conn //, _ = nats.Connect("nats://127.0.0.1:14222")

func Subscribe() {

}

func HandleResponse(data *P.FRMSMessage_Typologyresult) {
	natsConnection, _ = nats.Connect("nats://127.0.0.1:14222")
	log.Println("Connected to " + nats.DefaultURL)

	log.Println("Sending Message...")
	// Publish message on subject
	// message := json.Unmarshal([]byte(data), string)
	// message, _ := json.Marshal(data)

	message, _ := proto.Marshal(data)

	natsConnection.Publish("TPOUT", message)
	log.Println("Message Send.")
}
