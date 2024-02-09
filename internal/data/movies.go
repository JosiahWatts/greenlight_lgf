package data

import "time"

// omitempty will hide the fiels in the output
// if and only if they are empty
//
// adding a string will make the field a string in
// the json (works on int, uint, float or bool types)
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
