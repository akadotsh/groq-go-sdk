package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/akadotsh/groq-go/pkg/types"
)

type Groq struct {
	ApiKey string
	Model  types.GroqModel
}

func (groqInstance *Groq) Values() []types.GroqModel {
	return []types.GroqModel{
		types.Gemma2_9b_it,
		types.Gemma_7b_it,
		types.Llama_31_70b_versatile,
		types.Llama_31_8b_instant,
		types.Llama3_70b_8192,
		types.Llama3_8b_8192,
		types.Llama_guard_3_8b,
		types.Mixtral_8x7b_32768,
		types.Whisper_large_v3,
	}
}


func (g *Groq) Chat(message string) (*types.Response, error) {
	url := "https://api.groq.com/openai/v1/chat/completions"

	if g.ApiKey == "" {
		return nil, errors.New("API key not found")
	}

	constructBody := types.Body{
		Messages: []types.Messages{{
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

	var response types.Response

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, err
}
