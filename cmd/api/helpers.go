package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// Define an envelope type
type envelope map[string]interface{}

// Define a writeJSON() helper for sending responses. This takes the destination
// http.ResponseWriter, the Http status code to send, the data to encode to JSON and a
// header map containing any additional HTTP headers we want to include in the response
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	// Use the json.MarshalIndent() function so that whitespace is added to the encoded
	// JSON. Here we use no line prefix ("") and tab indents ("\t") for each element.
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return err
	}

	// append a newline to the JSON. This is just a small nicety to make it easier
	// to view in terminal apps
	js = append(js, '\n')

	// Note that it's OK if the provided header map is nil. Go doesn't throw an error
	// if you try to range over (or generally, read from) a nil map.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// add the content type header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
