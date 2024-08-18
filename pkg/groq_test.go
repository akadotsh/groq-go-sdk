package pkg

import (
	"reflect"
	"testing"
)

func TestChat(t *testing.T) {
	groq := Groq{
		ApiKey: "",
		Model:  Gemma2_9b_it,
	}

	_, err := groq.Chat("Explain the importance of fast language models")

	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	expectedErrorMsg := "API key not found"

	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMsg, err.Error())
	}

}

func TestGetAllModelValues(t *testing.T) {
	groq := Groq{
		ApiKey: "",
		Model:  Gemma2_9b_it,
	}

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
		t.Errorf("Expected error message '%s', but got '%s'", expectedValue, models)
	}
}
