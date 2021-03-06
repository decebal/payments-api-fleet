package errorHandler

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error" binding:"required"`
}

func OutputHTTPError(e string, w http.ResponseWriter, status int) {
	errResponse := errorResponse{e}
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(errResponse)
	if err != nil {
		panic(err)
	}
}
