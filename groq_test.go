package groq

import (
	"reflect"
	"testing"
)

func TestNoAPIKey(t *testing.T) {
	groq := New("")

	_, err := groq.Chat(Chat{
		Messages: []Message{
			{
				Role:    User,
				Content: "Explain the importance of fast language models",
			},
		},
		Model: Mixtral_8x7b_32768,
	})

	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	expectedErrorMsg := "API key not found"

	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMsg, err.Error())
	}
}

func TestGetAllModelValues(t *testing.T) {
	groq := New("dummy-api-key")

	models := groq.GetModels()

	expectedValue := []GroqModel{
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

	if !reflect.DeepEqual(expectedValue, models) {
		t.Errorf("Expected models %v, but got %v", expectedValue, models)
	}
}
