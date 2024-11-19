package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aryan1306/ai-commit-gen/internal"
	"github.com/aryan1306/ai-commit-gen/internal/config"

	"github.com/briandowns/spinner"
)

type openAiResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index     int    `json:"index"`
		Message   Message `json:"message"`
		Logprobs  string `json:"logprobs"`
		Finish    string `json:"finish_reason"`
		FinishRaw string `json:"finish_reason_raw"`
	} `json:"choices"`
	Usage struct {
		Prompt_tokens     int `json:"prompt_tokens"`
		Completion_tokens int `json:"completion_tokens"`
		Total_tokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func OpenAiClient(s *spinner.Spinner) {
	localConfig := config.LoadConfig()
	if localConfig.OpenAIKey == "" {
		log.Fatal("OpenAI API key not found in config file")
		os.Exit(1)
	}
	diffOut := internal.GenerateDiff()
	OpenAiRequest := RequestBody{
		Model: "gpt-4o",
		Messages: []Message{
			{
				Role:    "system",
				Content: SYSTEM_PROMPT,
			},
			{
				Role:    ROLE,
				Content: STARTER_PROMPT + "\n" + string(diffOut),
			},
		},
		Stream: false,
	}
	jsonBody, err := json.Marshal(OpenAiRequest)
	if err != nil {
		fmt.Println("Error marshalling JSON")
		os.Exit(1)
	}
	s.Start()
	req, httpErr := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if httpErr != nil {
		fmt.Println("Error creating HTTP request")
		s.Stop()
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ localConfig.OpenAIKey)
	client := &http.Client{}

	resp, respErr := client.Do(req)
	if respErr != nil {
		fmt.Println("Error sending HTTP request")
		s.Stop()
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.Stop()
		fmt.Println("Error: HTTP status code is not 200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		os.Exit(1)
	}
	var readableResponse openAiResponse
	if err := json.Unmarshal(body, &readableResponse); err != nil {
		log.Printf("Error unmarshalling response: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(readableResponse.Choices[0].Message.Content)
	s.Stop()
}