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
	// The main thing to note here is that only EXPORTED feilds can marshal so we have to use fields with Capital letter

	data, err := json.MarshalIndent(movies, " ", "  ")
	if err != nil {
		log.Fatalf("Json Marshalling failed : %s", err)
	}
	fmt.Printf("%s\n", data)

	//To decode JSON format to GO datastructure is called Unmarshaling.

	var newMovies []struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	newErr := json.Unmarshal(data, &newMovies)

	if newErr != nil {
		log.Fatalf("JSON unmarshalling failed: %s", newErr)
	} else {
		fmt.Println(newMovies) //[{Casablanca} {Scarface} {Milk}]
	}
}
