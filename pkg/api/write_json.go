package api

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if data == nil {
		w.Write([]byte(`{}`))
		return
	}

	result, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "ошибка формирования ответа json"}`))
		return
	}

	w.Write(result)
}
