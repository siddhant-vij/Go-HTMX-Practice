package main

import (
	"Click_To_Load/data"
	"Click_To_Load/templates"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageComponent := templates.Page()
		pageComponent.Render(r.Context(), w)
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		responseComponent := templates.DisplayResponse(data.ListItems)
		responseComponent.Render(r.Context(), w)
	})

	log.Println("Starting server on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
