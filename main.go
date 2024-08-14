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

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Body struct {
	Messages []Messages `json:"messages"`
	Model    GroqModel  `json:"model"`
}

type Choices struct {
	Index         int      `json:"index"`
	Message       Messages `json:"message"`
	Logprobs      any      `json:"logprobs"`
	Finish_Reason string   `json:"finish_reason"`
}

type Usage struct {
	Prompt_Tokens     float64 `json:"prompt_tokens"`
	Prompt_Time       float64 `json:"prompt_time"`
	Completion_Tokens int64   `json:"completion_tokens"`
	Completion_Time   float64 `json:"completion_time"`
	Total_Tokens      int64   `json:"total_tokens"`
	Total_Time        float64 `json:"total_time"`
}

type XGroq struct {
	ID string `json:"id"`
}

type Response struct {
	ID                 string    `json:"id"`
	Object             string    `json:"object"`
	Created            int       `json:"created"`
	Choices            []Choices `json:"choices"`
	Usage              Usage     `json:"usage"`
	System_Fingerprint string    `json:"system_fingerprint"`
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

func (groqInstance *Groq) Chat(message string) Response {
	url := "https://api.groq.com/openai/v1/chat/completions"

	constructBody := Body{
		Messages: []Messages{{
			Role:    "user",
			Content: message,
		},
		},
		Model: groqInstance.Model,
	}

	body, err := json.Marshal(constructBody)

	if err != nil {
		panic(err)
	}

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

	var resp Response
	derr := json.NewDecoder(res.Body).Decode(&resp)

	if derr != nil {
		panic(derr)
	}

	fmt.Println("resp", resp)

	return resp
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

	groq.Chat("Explain the importance of fast language models")

}
