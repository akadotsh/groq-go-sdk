package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
	Messages    []Messages `json:"messages"`
	Model       GroqModel  `json:"model"`
	Temperature float64    `json:"temperature"`
	Max_Tokens  float64    `json:"max_tokens"`
	Stream      bool       `json:"stream"`
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

func (g *Groq) Chat(message string) (*Response, error) {
	url := "https://api.groq.com/openai/v1/chat/completions"

	if g.ApiKey == "" {
		return nil, errors.New("API key not found")
	}

	constructBody := Body{
		Messages: []Messages{{
			Role:    "user",
			Content: message,
		},
		},
		Model:       g.Model,
		Temperature: 0.5,
		Stream:      false,
		Max_Tokens:  1024,
	}

	body, err := json.Marshal(constructBody)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	httpClient := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+g.ApiKey)

	if err != nil {
		fmt.Println("Error", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer req.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var response Response

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	fmt.Println("resp", response)

	return &response, err
}
