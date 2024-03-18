package models

// Request struct for the JSON request
type GPTRequest struct {
	Model       string       `json:"model"`
	Messages    []GPTMessage `json:"messages"`
	Temperature float64      `json:"temperature"`
}