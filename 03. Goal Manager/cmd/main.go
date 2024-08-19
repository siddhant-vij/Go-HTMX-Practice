package main

import (
	"Goal_Manager/templates"
	"log"
	"net/http"
	"time"
)

var courseGoals = []templates.CourseGoal{}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageComponent := templates.Page(courseGoals)
		pageComponent.Render(r.Context(), w)
	})

	http.HandleFunc("/addGoal", func(w http.ResponseWriter, r *http.Request) {
		goal := r.FormValue("goal")
		if goal != "" {
			newGoal := templates.CourseGoal{
				Index: time.Now().Format("20060102150405"),
				Goal:  goal,
			}
			courseGoals = append(courseGoals, newGoal)
			goalComponent := templates.DisplayGoal(newGoal.Index, newGoal.Goal)
			goalComponent.Render(r.Context(), w)
		}
	})

	http.HandleFunc("/removeGoal", func(w http.ResponseWriter, r *http.Request) {
		idx := r.FormValue("idx")
		if idx != "" {
			for i, goal := range courseGoals {
				if goal.Index == idx {
					courseGoals = append(courseGoals[:i], courseGoals[i+1:]...)
					break
				}
			}
		}
	})

	log.Println("Starting server on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
