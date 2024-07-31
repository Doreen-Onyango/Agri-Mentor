package main

import (
	"net/http"

	server "agri-mentor/agri"
)

func main() {
	// Serve static files from the /static directory
	http.Handle("/", http.FileServer(http.Dir("./templates")))

	// http.HandleFunc("/", server.RenderTemplate) // Serve the HTML file

	http.HandleFunc("/query", server.HandleSendMessage) // Handle the input submission

	println("Server started at port: 8081")
	http.ListenAndServe(":8081", nil) // Start the server
}
