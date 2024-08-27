package groq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	apiURL = "https://api.groq.com/openai/v1/chat/completions"
)

type Groq struct {
	apiKey string
}

func New(apiKey string) *Groq {
	return &Groq{
		apiKey: apiKey,
	}
}

func (groqInstance *Groq) GetModels() []GroqModel {
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

func (g *Groq) Chat(chat Chat) (*Response, error) {

	if g.apiKey == "" {
		return nil, errors.New("API key not found")
	}

	message := chat.Messages
	model := chat.Model

	constructBody := Body{
		Messages:    message,
		Model:       model,
		Temperature: 0.5,
		Stream:      false,
		Max_Tokens:  1024,
	}

	body, err := json.Marshal(constructBody)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+g.apiKey)

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

	return &response, err
}