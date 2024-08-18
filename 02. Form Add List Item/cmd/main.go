package main

import (
	"Form_Add_List_Item/data"
	"Form_Add_List_Item/templates"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageComponent := templates.Page(data.ListItems)
		pageComponent.Render(r.Context(), w)
	})

	http.HandleFunc("/addNote", func(w http.ResponseWriter, r *http.Request) {
		note := r.FormValue("note")
		if note != "" {
			listItem := templates.DisplayListResponse(note)
			listItem.Render(r.Context(), w)

			formClear := templates.FormResponseOOB()
			formClear.Render(r.Context(), w)
		}
	})

	log.Println("Starting server on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
