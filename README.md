# Groq Go SDK

## Installation

```bash
go get github.com/akadotsh/groq-go-sdk
```

## Usage

1. Get an API key from [https://console.groq.com/keys](https://console.groq.com/keys)
2. Set the API key as an environment variable named `groq_api_key`.

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akadotsh/groq-go-sdk"
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
