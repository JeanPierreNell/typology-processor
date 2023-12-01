package main

import (
	"runtime"

	db "typology-processor/db-lib"

	"github.com/joho/godotenv"
)

// Client

func main() {

	err := godotenv.Load()

	if err != nil {
		println(err)
	}

	// Init Arango DB
	db.InitDatabases()

	// Init NATS Subscription
	Subscribe(HandleTransaction)

	// Keep the connection alive
	runtime.Goexit()
}
