package main

import (
	"fmt"
	"net/http"

	server "agric/server"
)

func main() {
	r := http.NewServeMux()
	//	fileserver := http.FileServer(http.Dir("/templates"))
	r.HandleFunc("/", server.Handl)
	r.HandleFunc("/signin", server.RenderTemplate)
	r.HandleFunc("/chatbot", server.Handl)

	server := http.Server{
		Addr:    ":8081",
		Handler: r,
	}
	fmt.Println("Server listening on port http://localhost:8081")
	server.ListenAndServe()
}
