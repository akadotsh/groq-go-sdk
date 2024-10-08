> ### IMPORTANT: This is an unofficial, community-developed Groq client library for Go. It is not affiliated with, officially maintained, or endorsed by Groq Inc.

# Groq Go SDK

### Go client library for interacting with the Groq API.

## Requirements
- Go 1.22.4 or higher
- [godotenv package](https://github.com/joho/godotenv) (optional, for loading environment variables from a `.env` file)


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
	groqAPIKey := os.Getenv("GROQ_API_KEY")
	if groqAPIKey == "" {
		log.Fatal("GROQ_API_KEY not set in environment variables")
	}

	// Initialize the Groq client
	client := groq.New(groqAPIKey)
	if client == nil {
		log.Fatal("Failed to create Groq client")
	}

	// Prepare the chat request
	chat := groq.Chat{
		Messages: []groq.Message{
			{
				Role:    groq.User,
				Content: "Explain the importance of fast language models",
			},
		},
		Model: groq.Mixtral_8x7b_32768,
	}

	// Send a chat request
	response, err := client.Chat(chat)
	if err != nil {
		log.Fatalf("Error calling Chat: %v", err)
	}

	// Print the response
	fmt.Println(response)
}
```

# License
This project is licensed under the MIT License - [see the LICENSE file for details](https://github.com/akadotsh/groq-go-sdk?tab=MIT-1-ov-file)
