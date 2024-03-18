package models

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ChatGPTClient struct {
	Organization string
	APIKey       string
	BaseURL      string
}

func NewChatGPTClient(org string, apiKey string) *ChatGPTClient {
	return &ChatGPTClient{
		Organization: org,
		APIKey:       apiKey,
		BaseURL:      "https://api.openai.com/v1/chat",
	}
}

func (c *ChatGPTClient) SetQuestion() string {
	fmt.Println("Enter your question: ")
	reader := bufio.NewReader(os.Stdin)
	var question, _ = reader.ReadString('\n')
	return question
}

func (c *ChatGPTClient) SendRequest(message string) (string, error) {

	requestBody, err := json.Marshal(GPTRequest{
		Model:       "gpt-3.5-turbo",
		Messages:    []GPTMessage{{Role: "user", Content: message}},
		Temperature: 0.7,
	})
	fmt.Println("-\nRequest :")
	fmt.Println(string(requestBody))
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/completions", c.BaseURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("OpenAI-Organization", c.Organization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("-\nResponse :")
	fmt.Println(string(responseBody))

	var completionResponse CompletionResponse
	err = json.Unmarshal(responseBody, &completionResponse)
	if err != nil {
		return "", err
	}

	if len(completionResponse.Choices) > 0 {
		return completionResponse.Choices[0].Text, nil
	}

	return "", nil
}