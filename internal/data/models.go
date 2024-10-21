package data

import (
	"database/sql"
	"errors"
)

// define a custom ErrRecordNotFound error. We'll return this
// from our Get() method when looking up a movie that DNE
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// create a models struct which wraps the movie model
type Models struct {
	Movies MovieModel
}

// for ease of use, we also add a New() method which returns
// a Models struct containing the initialized movie model
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
