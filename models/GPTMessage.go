package models

// GPTMessage struct representing messages in the request
type GPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}