package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Initialize the OpenAI client with your API key
	apiKey := os.Getenv("Authorization: Bearer sk-None-2HcyFwqVMe6kuaKH3BOWT3BlbkFJz5ncWmGjub4DoUyE3vdP") // Ensure your API key is set in the environment
	client := openai.NewClient(apiKey)
	// fmt.Println("API Key:", apiKey)

	// Define the region context
	region := "Kisumu"
	contextPrompt := fmt.Sprintf("You are a chatbot that can only provide information about %s.", region)

	// Start the conversation loop
	for {
		var userInput string
		fmt.Print("You: ")
		fmt.Scanln(&userInput)

		// Create a chat completion request
		resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: contextPrompt},
				{Role: openai.ChatMessageRoleUser, Content: userInput},
			},
		})
		if err != nil {
			fmt.Printf("Error generating response: %v\n", err)
			continue
		}

		// Output the AI's response
		fmt.Printf("AI: %s\n", resp.Choices[0].Message.Content)
	}
}
