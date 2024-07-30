package main

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

type Chatbot struct {
	client *openai.Client
}

func NewChatbot() *Chatbot {
	client := openai.NewClient(os.Getenv("Authorization: Bearer"))
	return &Chatbot{client: client}
}

func (c *Chatbot) GenerateResponse(query string) (string, error) {
	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are an AI assistant for Kenyan farmers. Provide advice on crop management, pest control, and market trends.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
