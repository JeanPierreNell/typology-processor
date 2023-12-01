package main

import (
	"log"
	"os"

	P "typology-processor/proto"

	"google.golang.org/protobuf/proto"

	"github.com/nats-io/nats.go"
)

var natsConnection *nats.Conn

func Subscribe() {

}

func HandleResponse(data *P.FRMSMessage) {
	natsConnection, _ = nats.Connect(os.Getenv("SERVER_URL"))
	log.Println("Connected to " + nats.DefaultURL)

	log.Println("Sending Message...")
	// Publish message on subject
	// message := json.Unmarshal([]byte(data), string)
	// message, _ := json.Marshal(data)

	message, _ := proto.Marshal(data)

	natsConnection.Publish(os.Getenv("PRODUCER_STREAM"), message)
	log.Println("Message Send.")
}
