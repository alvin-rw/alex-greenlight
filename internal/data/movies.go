package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`    // movie release year
	Runtime   int32     `json:"runtime"` // movie runtime in minutes
	Genres    []string  `json:"genres"`
	Version   int32     `json:"version"` // starts at 1 and will be incremented each time the movie information is updated
}
