package main

import (
	"log"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return
	}
	log.Println("Loaded .env file ")
}

func main() {
	loadEnv()
}
