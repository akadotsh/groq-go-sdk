package pkg

import (
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
