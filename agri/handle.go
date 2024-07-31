package agri

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Message struct {
	Content string `json:"content"`
}

type PageData struct {
	Messages []Message
}

var messages []Message

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	// Parse for data from the client
	// 	if err := r.ParseForm(); err != nil {
	// 		http.Error(w, fmt.Sprintf("ParseForm() %v", err), http.StatusBadRequest)
	// 		return
	// 	}
	// 	t, err := template.ParseFiles("templates/index.html")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}

	// 	err = t.Execute(w, struct{ Messages []string }{Messages: messages})

	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// }

	tmpl, err := template.ParseFiles(filepath.Join("templates","index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Messages: messages,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
