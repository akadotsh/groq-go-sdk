package main

import (
	"fmt"
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

	response,err:= groq.Chat("Explain the importance of fast language models")

	if err != nil{
		fmt.Println("Error",err)
	}

	fmt.Println(response)

}
