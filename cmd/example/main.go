package main

import (
	"log"
	"os"

	groq "github.com/akadotsh/groq-go/pkg"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	groq_api_key := os.Getenv("groq_api_key")

	groq := groq.Groq{
		ApiKey: groq_api_key,
		Model:  groq.Mixtral_8x7b_32768,
	}

	groq.Chat("Explain the anatomy of human life")

}
