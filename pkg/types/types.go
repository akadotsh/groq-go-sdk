package types


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

