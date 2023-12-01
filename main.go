package main

import (
	"log"
	"runtime"

	db "typology-processor/db-lib"
	P "typology-processor/proto"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

const subject = "TP"

// Client

func main() {

	// Create server connection
	//nats.DefaultURL
	natsConnection, _ := nats.Connect("nats://127.0.0.1:14222")
	log.Println("Connected to " + "nats://127.0.0.1:14222")

	db.InitDatabases()

	// Subscribe to subject & Handle incoming messages.
	natsConnection.Subscribe(subject, func(msg *nats.Msg) {
		log.Println("Recieved Message. Processing...")

		message := &P.FRMSMessage{}
		err := proto.Unmarshal(msg.Data, message)

		if err != nil {
			log.Fatal("Could not unmarshal Protobuff object.")
		}

		// HandleTransaction(string(msg.Data))
		HandleTransaction(message)

		log.Println("Message Resolved.")
	})

	// Keep the connection alive
	runtime.Goexit()
}
