package main

import (
	"go-microservices/internal/database"
	"go-microservices/internal/server"
	"log"
)

func main() {
	//setup db
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to init db client: %s", err)
	}

	//setup server
	srv := server.NewEchoServer(db)
	err = srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
