package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type GroqModel string

const (
	Gemma2_9b_it           GroqModel = "gemma2-9b-it"
	Gemma_7b_it            GroqModel = "gemma-7b-it"
	Llama_31_70b_versatile GroqModel = "llama_31_70b_versatile"
	Llama_31_8b_instant    GroqModel = "llama-3.1-70b-versatile"
	Llama3_70b_8192        GroqModel = "llama3-70b-8192"
	Llama3_8b_8192         GroqModel = "llama3-8b-8192"
	Llama_guard_3_8b       GroqModel = "llama-guard-3-8b"
	Mixtral_8x7b_32768     GroqModel = "mixtral-8x7b-32768"
	Whisper_large_v3       GroqModel = "whisper-large-v3"
)

type Groq struct {
	ApiKey string
	Model  GroqModel
}

func (groqInstance *Groq) Values() []GroqModel {
	return []GroqModel{
		Gemma2_9b_it,
		Gemma_7b_it,
		Llama_31_70b_versatile,
		Llama_31_8b_instant,
		Llama3_70b_8192,
		Llama3_8b_8192,
		Llama_guard_3_8b,
		Mixtral_8x7b_32768,
		Whisper_large_v3,
	}
}

func (groqInstance *Groq) Chat() {
	url := "https://api.groq.com/openai/v1/chat/completions"

	body := []byte(`{
    "messages":[{
          "role": "user",
          "content": "Explain the importance of fast language models"
        }],
      "model":"mixtral-8x7b-32768"  
    }`)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+groqInstance.ApiKey)

	if err != nil {
		fmt.Println("Error", err)
		log.Panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	var resp map[string]interface{}
	derr := json.NewDecoder(res.Body).Decode(&resp)

	if derr != nil {
		panic(derr)
	}

	fmt.Println("resp", resp)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	groq_api_key := os.Getenv("groq_api_key")

	groq := Groq{
		ApiKey: groq_api_key,
		Model:  Mixtral_8x7b_32768,
	}

	groq.Chat()

}
