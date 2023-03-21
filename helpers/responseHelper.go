package helpers

import (
	"encoding/json"
	"net/http"
)

func HttpResponse(success bool, message string, body []byte, httpStatus int, w http.ResponseWriter) {

	data := map[string]interface{}{
		"success": success,
		"message": message,
		"data":    body,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}
