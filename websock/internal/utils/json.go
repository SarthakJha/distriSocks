package utils

import (
	"encoding/json"
	"io"
)

func ToJSON(data interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(data)
}

func FromJSON(payload interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(payload)
}
