package util

import (
	"encoding/json"
	"net/http"
)

const ContentTypeJson = "application/json"

// Render interface for HTTP response formats.
type Render interface {
	// Error response rendering.
	Error(statusCode int, message string) error
	// JSON response rendering with a 200 status.
	JSON(data interface{}) error
}

// HttpRender encapsulates a HTTP response stream.
type HttpRender struct {
	http.ResponseWriter
}

// Renderer constructor.
func Renderer(w http.ResponseWriter) Render {
	return HttpRender{w}
}

// Error response rendering.
func (r HttpRender) Error(statusCode int, message string) error {
	response, _ := json.Marshal(map[string]string{"message": message})

	r.Header().Set("Content-Type", ContentTypeJson)
	r.WriteHeader(statusCode)

	_, e := r.Write(response)
	if e != nil {
		return e
	}

	return nil
}

// JSON response rendering with a 200 status.
func (r HttpRender) JSON(data interface{}) error {
	response, err := json.Marshal(data)
	if err != nil {
		return r.Error(http.StatusInternalServerError, err.Error())
	}

	r.Header().Set("Content-Type", ContentTypeJson)
	r.WriteHeader(http.StatusOK)

	_, e := r.Write(response)
	if e != nil {
		return e
	}

	return nil
}
