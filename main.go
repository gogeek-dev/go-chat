package main

import (
	"go-realtime-chat/routes"
	"log"
)

func main() {

	r := routes.SetupRoutes()

	log.Println("listening on http://localhost:8040")
	r.Run(":8040")

}
