package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const ApplicationJson = "application/json"

type Request struct {
	*http.Request
}

// Decode request to target interface type.
func (r Request) Decode(decodedType interface{}) (err error) {
	body := r.Body

	if r.Body != nil {
		defer r.Body.Close()
	}

	if r.Header.Get("Content-Type") != ApplicationJson {
		return errors.New("Unsupported Content-Type")
	}

	if body != nil {
		d := json.NewDecoder(body)
		d.DisallowUnknownFields() // Strict mode (i.e. error if unknown field is encountered)

		if err := d.Decode(&decodedType); err != nil {
			msg := err.Error()

			// Cleaner JSON unmarshal error message
			if strings.HasPrefix(msg, "json: cannot unmarshal") {
				v, ok := err.(*json.UnmarshalTypeError)

				if ok {
					msg = strings.ToLower(v.Field) + ": invalid type."
				}
			}

			return errors.New(msg)
		}
	}

	return err
}
