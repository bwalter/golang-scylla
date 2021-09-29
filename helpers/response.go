package helpers

import (
	"encoding/json"
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
func RespondWithError(w http.ResponseWriter, code int, message string) error {
	return RespondWithJSON(w, code, ErrorObject{Error: message})
}

// => (code: <code> body: payload JSON)
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

	return nil
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
