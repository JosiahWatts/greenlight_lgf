package data

import (
	"fmt"
	"strconv"
)

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
