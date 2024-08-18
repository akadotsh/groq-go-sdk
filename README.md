# Groq Go SDK

## Usage

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akadotsh/groq-go-sdk" // Replace with your actual import path
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the API key from environment variables
	groqAPIKey := os.Getenv("groq_api_key")

	// Initialize the Groq client
	groq := groq.Groq{
		ApiKey: groqAPIKey,
		Model:  groq.Mixtral_8x7b_32768,
	}

	// Send a chat request
	response, err := groq.Chat("Explain the importance of fast language models")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}


	fmt.Println(response)
}
```
