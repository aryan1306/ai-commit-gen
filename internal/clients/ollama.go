package clients

import (
	"ai-commits/internal"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/briandowns/spinner"
)

type RequestBody struct {
	Model string `json:"model"`
	Messages []Message `json:"messages"`
	Stream bool `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Stream  bool   `json:"stream"`
}

type OllamaResponse struct {
	Message Message `json:"message"`
	Done    bool    `json:"done"`
}

const (
	MODEL = "gemma2:latest"
	ROLE = "user"
	SYSTEM_PROMPT = `You are a helpful assistant that explains Git changes in a concise way.
	Focus only on the most significant changes and their direct impact.
	When answering specific questions, address them directly and precisely.
	Keep explanations brief but informative and don't ask for further explanations.
	Use markdown for clarity.`
	STARTER_PROMPT = `Generate a concise git commit message written in present tense for the following code diff with the given specifications below:
	The output response must be in format:<type>(<optional scope>): <commit message>
	Choose a type from the following list:
	- feat: A new feature
	- fix: A bug fix
	- docs: Documentation only changes
	- style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
	- refactor: A code change that neither fixes a bug nor adds a feature
	- perf: A code change that improves performance
	- test: Adding missing tests or correcting existing tests
	- chore: Changes to the build process or auxiliary tools and libraries such as documentation generation
	Focus on being accurate and concise.
	Commit message must be a maximum of 72 characters.
	Exclude anything unnecessary such as translation. Your entire response will be passed directly into git commit.
	Code diff: `
)

func OllamaClient(s *spinner.Spinner, modelFlag string) {
	diffOut := internal.GenerateDiff()

	llamaRequest := RequestBody{
		Model: modelFlag,
		Messages: []Message{
			{
				Role:    "user",
				Content: SYSTEM_PROMPT + "\n" + STARTER_PROMPT + "\n" + string(diffOut),
			},
		},
		Stream: true,
	}
	jsonBody, err := json.Marshal(llamaRequest)
	if err != nil {
		fmt.Println("Error marshalling JSON")
		os.Exit(1)
	}
	s.Start()
	req, httpErr := http.NewRequest("POST", "http://localhost:11434/api/chat", bytes.NewBuffer(jsonBody))
	if httpErr != nil {
		fmt.Println("Error creating HTTP request")
		s.Stop()
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
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
	decoder := json.NewDecoder(resp.Body)
	for {
		var response OllamaResponse
		if err := decoder.Decode(&response); err != nil {
			if err == io.EOF {
				log.Println("End of response stream")
				break
			}
			log.Printf("Error decoding response: %v\n", err)
			return
		}

		fmt.Print(response.Message.Content)
		if response.Done {
			log.Println("\nResponse complete")
			break
		}
		s.Stop()
	}
}