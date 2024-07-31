package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func Handl(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Error 404: PAGE NOT FOUND", http.StatusNotFound)
			fmt.Println(err)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} else if r.URL.Path == "/styles.css" {
			http.ServeFile(w, r, "templates/styles.css")
		} else if r.URL.Path == "/login.html" {
			http.ServeFile(w, r, "templates/login.html")
		} else if r.URL.Path == "/about.html" {
			http.ServeFile(w, r, "templates/about.html")
		}
	}
}
