package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	// Generate a string containing the movie runtime
	// in the required format.
	jsonValue := fmt.Sprintf("%d mins", r)

	// use the strconv.Quote() function on the string to
	// wrap it in doible quotes. Needs quotes to be valid
	quotedJSONValue := strconv.Quote(jsonValue)

	// convert the quoted string value to a byte slice
	return []byte(quotedJSONValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// we expect the incoming json will be a string in the format
	// <runtime> mins, and the first we need to do is remove the surrounding
	// quotes from the string.
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// split parts to insolate the part containing the number
	parts := strings.Split(unquotedJSONValue, " ")

	// Sanity check the parts of the string to make sure it was in the expected format.
	// If it isn't, we return the ErrInvalidRuntimeFormat error again.
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	// otherwise, parse the string containing the number into an int32. Again,
	// if this fails return the error
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
