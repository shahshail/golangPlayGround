package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"aaa", "bbb", "ccc"}},
		{Title: "Scarface", Year: 1988, Color: true, Actors: []string{"ddd", "eee", "fff"}},
		{Title: "Milk", Year: 2017, Color: true, Actors: []string{"ggg", "hhh", "iii"}},
	}

	// Converting GO Datastructure in JOSON called Marshalling
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Json Marshalling failed : %s", err)
	}

	fmt.Printf("%s\n", data)
}
