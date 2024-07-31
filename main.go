package main

import (
	"net/http"

	"agri-mentor/server"
)

func main() {
	// Serve static files from the /static directory
    http.Handle("/", http.FileServer(http.Dir("./templates")))

	// http.HandleFunc("/", server.RenderTemplate) // Serve the HTML file
	
	http.HandleFunc("/query", server.HandleSendMessage) // Handle the input submission

	println("Server started at port: 8080")
	http.ListenAndServe(":8080", nil) // Start the server
}
