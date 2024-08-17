package pkg

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
