package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	router := gin.Default()

	router.POST("/api/query", app.HandleQuery)
	router.POST("/api/pest-report", app.HandlePestReport)
	router.GET("/api/pest-occurrences", app.HandleGetPestOccurrences)
	router.POST("/api/feedback", app.HandleFeedback)

	router.Run(":8080")
}
