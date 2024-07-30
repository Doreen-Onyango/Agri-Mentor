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
		} else if r.URL.Path == "/bot.html" {
			http.ServeFile(w, r, "templates/bot.html")
		}

		// } else if r.URL.Path == "/styles.css" {
		// 	cssPath := filepath.Join("../", "templates", "styles.css")
		// 	fmt.Println(cssPath)
		// 	http.ServeFile(w, r, cssPath)
		// } else if r.URL.Path == "/bot.html" {
		// 	about := filepath.Join("../", "templates", "bot.html")
		// 	http.ServeFile(w, r, about)
		// } else if r.URL.Path == "/error.css" {
		// 	about := filepath.Join("../", "templates", "error.css")
		// 	http.ServeFile(w, r, about)
		// } else {
		// 	path := filepath.Join("/", "templates", "error.html")
		// 	tmpl, err := template.ParseFiles(path)
		// 	if err != nil {
		// 		w.WriteHeader(http.StatusNotFound)
		// 	}
		// 	tmpl.Execute(w, nil)
		// }
	}
}
