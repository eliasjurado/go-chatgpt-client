package main

import (
	"fmt"

	"chat-client/models"
	"chat-client/resources"
)

func main() {
	client := models.NewChatGPTClient(resources.ORG, resources.APIKEY)

	question := client.SetQuestion()

	response, err := client.SendRequest(question)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("ChatGPT Response:", response)
}
