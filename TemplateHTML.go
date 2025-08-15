package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template.html"))

	data := struct {
		Title   string
		Message string
		Name    string
	}{
		Title:   "Welcome to My Website",
		Message: "Hello, World!",
		Name:    "Pradeep",
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
