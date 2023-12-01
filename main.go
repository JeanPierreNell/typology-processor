package main

import (
	"log"
	"os"
	"runtime"

	db "typology-processor/db-lib"
	P "typology-processor/proto"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

// Client

func main() {

	godotenv.Load()
	subject := os.Getenv("CONSUMER_STREAM")

	// Create server connection
	//nats.DefaultURL
	//nats://127.0.0.1:14222
	natsConnection, natsError := nats.Connect(os.Getenv("SERVER_URL"))

	if natsError != nil {
		log.Fatal("Could not connect to NATS")
		panic("Could not connect to NATS")
	}

	log.Println("Connected to NATS on: " + os.Getenv("SERVER_URL"))

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
