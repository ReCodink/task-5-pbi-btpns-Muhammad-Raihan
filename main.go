package main

import (
	"log"
	"os"

	db "github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/database"
	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/router"
	"github.com/joho/godotenv"
)

func main() {
	db.Init()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	route := router.Router()

	go func() {
		if err := route.Run(":" + port); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	log.Printf("Server started on port %s", port)

	select {}
}
