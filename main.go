package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token string
	BotName string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token = os.Getenv("TOKEN")
	BotName = os.Getenv("CLIENT_ID")
}

func main() {
	fmt.Println("Hello, DisGo.")
	fmt.Printf("Token: %s\n", Token)
	fmt.Printf("BotName: %s\n", BotName)
}