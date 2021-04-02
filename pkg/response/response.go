package response

import (
	"encoding/json"
	"net/http"
)


type ErrorMessage struct {
	Message string `json:"message"`
}

type Map map[string]interface{}

func JSON(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) error {

	w.Header().Set("content-type", "application/json")

	if data == nil {
		w.WriteHeader(statusCode)
		return nil
	}

	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	
	w.WriteHeader(statusCode)
	w.Write(j)

	return nil
}

func HTTPError(w http.ResponseWriter, r *http.Request, statusCode int, message string) error {
	msg := ErrorMessage{
		Message: message,
	}

	return JSON(w, r, statusCode, msg)
}