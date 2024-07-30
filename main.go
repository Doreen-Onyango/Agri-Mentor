package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"agri-mentor/chatbot"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	apiKey, found := os.LookupEnv("GEMINI_API_KEY")
	if !found {
		log.Fatal("Environment variable GEMINI_API_KEY not set\n")
	}
	option := option.WithAPIKey(apiKey)

	client, err := genai.NewClient(ctx, option)
	if err != nil {
		log.Fatalf("Error creating client: %v\n", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	model.SetTemperature(0.1)
	model.SetTopK(30)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(1000)
	model.ResponseMIMEType = "text/plain"

	// model.SafetySettings = Adjust safety settings
	// See https://ai.google.dev/gemini-api/docs/safety-settings

	prompt := "harvest time for basmat rice"
	answer, err := chatbot.ProcessQuery(prompt, "resource/data.csv")
	if err != nil {
		fmt.Println("Error:", err)
	} // else {
	// 	fmt.Println("Result:", answer)
	// }

	parts := []genai.Part{
		genai.Text("input: Respond only with the information provided but be wordy: " + prompt),
		genai.Text("output: " + answer),
	}

	resp, err := model.GenerateContent(ctx, parts...)
	if err != nil {
		log.Fatalf("Error sending message: %v\n", err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part)
	}
}
