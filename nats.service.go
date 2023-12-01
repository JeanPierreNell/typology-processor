package main

import (
	"log"
	"os"

	P "typology-processor/proto"

	"google.golang.org/protobuf/proto"

	"github.com/nats-io/nats.go"
)

type natsFunction func(*P.FRMSMessage)

var natsConnection *nats.Conn

func Subscribe(function natsFunction) {

	subject := os.Getenv("CONSUMER_STREAM")
	serverUrl := os.Getenv("SERVER_URL")
	var natsError error

	natsConnection, natsError = nats.Connect(serverUrl)

	if natsError != nil {
		log.Fatal("Could not connect to NATS")
		panic("Could not connect to NATS")
	}

	log.Println("Connected to NATS on: " + serverUrl)

	natsConnection.Subscribe(subject, func(msg *nats.Msg) {
		log.Println("Recieved Message. Processing...")

		message := &P.FRMSMessage{}
		err := proto.Unmarshal(msg.Data, message)

		if err != nil {
			log.Fatal("Could not unmarshal Protobuff object.")
		}

		function(message)
		log.Println("Message Resolved.")
	})
}

func HandleResponse(data *P.FRMSMessage) {
	producerStream := os.Getenv("PRODUCER_STREAM")

	log.Println("Sending Message...")
	message, _ := proto.Marshal(data)

	err := natsConnection.Publish(producerStream, message)

	if err != nil {
		log.Println("Failed to send message to NATS" + err.Error())
	} else {
		log.Println("Message Send.")
	}
}
