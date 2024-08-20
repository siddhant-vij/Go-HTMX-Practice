package utils

import (
	"encoding/json"
	"log"
	"os"
)

type Location struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Image struct {
		Src string `json:"src"`
		Alt string `json:"alt"`
	} `json:"image"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func ReadJsonFile(filePath string) []Location {
	var locations []Location

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&locations)
	if err != nil {
		log.Fatal(err)
	}

	return locations
}
