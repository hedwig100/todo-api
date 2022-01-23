package route

import (
	"encoding/json"
	"net/http"
)

type ErrorJson struct {
	ErrorMessage string `json:"errorMessage"`
}

func sendErrorMessage(writer http.ResponseWriter, message string) {
	body, _ := json.Marshal(ErrorJson{ErrorMessage: message})
	writer.WriteHeader(500)
	writer.Write(body)
}
