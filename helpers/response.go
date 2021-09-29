package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

type ErrorObject struct {
	Error string `json:"error"`
}

func NewErrorObject(error string) ErrorObject {
	return ErrorObject{Error: error}
}

// => (code: <code> body: error JSON)
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, ErrorObject{Error: message})
}

// => (code: <code> body: payload JSON)
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("ERROR: Could not respond: %v\n", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err = w.Write(response); err != nil {
		fmt.Printf("ERROR: Could not respond: %v\n", err.Error())
		return
	}
}

func DecodeJSONBody(body io.Reader, v interface{}) error {
	if err := json.NewDecoder(body).Decode(v); err != nil {
		return err
	}

	// Do not validate maps
	if _, ok := v.(*map[string]string); ok {
		return nil
	}

	validate := validator.New()
	return validate.Struct(v)
}
