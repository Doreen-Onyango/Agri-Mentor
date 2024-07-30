package main

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

type VectorEmbedding struct {
	client *openai.Client
}

func NewVectorEmbedding() *VectorEmbedding {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	return &VectorEmbedding{client: client}
}

func (ve *VectorEmbedding) GenerateEmbedding(text string) ([]float32, error) {
	resp, err := ve.client.CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Model: openai.AdaEmbeddingV2,
			Input: []string{text},
		},
	)
	if err != nil {
		return nil, err
	}

	return resp.Data[0].Embedding, nil
}
