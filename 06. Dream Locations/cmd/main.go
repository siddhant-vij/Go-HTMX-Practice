package main

import (
	"Dream_Locations/templates"
	"Dream_Locations/utils"
	"crypto/rand"
	"log"
	"math/big"
	"net/http"
)

var locations = utils.ReadJsonFile("data/locations.json")
var dreamLocations = []utils.Location{}
var suggestedLocations = []utils.Location{}

func chooseRandomLocations(n int) []utils.Location {
	var randomLocations []utils.Location
	for i := 0; i < n; i++ {
		randomIndexBigInt, _ := rand.Int(rand.Reader, big.NewInt(int64(len(locations))))
		randomIndex := int(randomIndexBigInt.Int64())
		randomLocations = append(randomLocations, locations[randomIndex])
		locations = append(locations[:randomIndex], locations[randomIndex+1:]...)
	}
	return randomLocations
}

func main() {
	suggestedLocations = chooseRandomLocations(4)

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageComponent := templates.Page(locations, dreamLocations, suggestedLocations)
		pageComponent.Render(r.Context(), w)
	})

	http.HandleFunc("/removeItem", func(w http.ResponseWriter, r *http.Request) {
		locationId := r.FormValue("item")
		if locationId != "" {
			curList := r.FormValue("curList")
			if curList == "available" {
				for i, location := range locations {
					if location.Id == locationId {
						locations = append(locations[:i], locations[i+1:]...)
						dreamLocations = append(dreamLocations, location)
						listItemUpd := templates.UpdateLocation("dream", &location)
						listItemUpd.Render(r.Context(), w)
						break
					}
				}
			} else if curList == "dream" {
				for i, location := range dreamLocations {
					if location.Id == locationId {
						dreamLocations = append(dreamLocations[:i], dreamLocations[i+1:]...)
						locations = append(locations, location)
						listItemUpd := templates.UpdateLocation("available", &location)
						listItemUpd.Render(r.Context(), w)
						break
					}
				}
			} else if curList == "suggested" {
				for i, location := range suggestedLocations {
					if location.Id == locationId {
						suggestedLocations = append(suggestedLocations[:i], suggestedLocations[i+1:]...)
						dreamLocations = append(dreamLocations, location)
						listItemUpd := templates.UpdateLocation("dream", &location)
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
