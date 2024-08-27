package main

import (
	"fmt"
	"log"
	"os"

	groqsdk "github.com/akadotsh/groq-go-sdk"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	groq_api_key := os.Getenv("groq_api_key")

	groq := groqsdk.New(groq_api_key)

	response, err := groq.Chat(groqsdk.Chat{
		Messages: []groqsdk.Message{
			{
				Role:    groqsdk.User,
				Content: "Explain the importance of fast language models",
			},
		},
		Model: groqsdk.Gemma2_9b_it,
	})

	if err != nil {
		log.Fatalf("Error calling Chat: %v", err)
	}

	
	fmt.Println(response.Choices[0].Message)

}
