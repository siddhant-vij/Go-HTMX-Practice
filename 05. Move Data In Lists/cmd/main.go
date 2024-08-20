package main

import (
	"Move_Data_In_Lists/data"
	"Move_Data_In_Lists/templates"
	"log"
	"net/http"
)

var leftList = data.ListItems
var rightList = []string{}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageComponent := templates.Page(leftList, rightList)
		pageComponent.Render(r.Context(), w)
	})

	http.HandleFunc("/removeItem", func(w http.ResponseWriter, r *http.Request) {
		item := r.FormValue("item")
		if item != "" {
			curList := r.FormValue("curList")
			if curList == "left" {
				for i, listItem := range leftList {
					if listItem == item {
						leftList = append(leftList[:i], leftList[i+1:]...)
						rightList = append(rightList, item)
						listItemUpd := templates.UpdateListItem(item, "right")
						listItemUpd.Render(r.Context(), w)
						break
					}
				}
			} else if curList == "right" {
				for i, listItem := range rightList {
					if listItem == item {
						rightList = append(rightList[:i], rightList[i+1:]...)
						leftList = append(leftList, item)
						listItemUpd := templates.UpdateListItem(item, "left")
						listItemUpd.Render(r.Context(), w)
						break
					}
				}
			} else {
				return
			}
		}
	})

	log.Println("Starting server on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
