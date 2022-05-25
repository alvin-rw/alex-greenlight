package data

import "time"

type Movie struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Year      int32 // movie release year
	Runtime   int32 // movie runtime in minutes
	Genres    []string
	Version   int32 // starts at 1 and will be incremented each time the movie information is updated
}
