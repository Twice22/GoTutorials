package main 

import (
	"encoding/json"
	"fmt"
	"log"
)

// `json:"released"` is an example of a field tag
// that is associated at compile time with the field of a struct.
// so the rendered json file will have released: instead of Year

// the first tag specify the JSON name, the other specify options
// NOTE: only exported fields are marshaled, which is why we choose
// capitalized names for all the Go field names (see p101)
type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
	Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
	Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
	Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	{
		//!+Marshal
		// print JSON
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
		//!-Marshal
	}

	{
		//!+MarshalIndent
		// pretty print the JSON
		data, err := json.MarshalIndent(movies, "", "    ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
		//!-MarshalIndent

		//!+Unmarshal
		// Convert JSON to slice of struct{ Title string}
		// as we only specified Title, only the Title will be
		// retrieve and save in the titles slice
		var titles []struct{ Title string }
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
		//!-Unmarshal
	}
}