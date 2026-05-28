package helper

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func ReadFromRequestBody(request *http.Request, result any) error {
	decoder := json.NewDecoder(request.Body)
	return decoder.Decode(result)
}

func WriteToResponseBody(writer http.ResponseWriter, statusCode int, response any) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	encoder := json.NewEncoder(writer)
	return encoder.Encode(response)
}

func WriteJSON(log *logrus.Logger, w http.ResponseWriter, status int, response any) {
	if err := WriteToResponseBody(w, status, response); err != nil {
		log.Errorf("failed to write response: %v", err)
	}
}
