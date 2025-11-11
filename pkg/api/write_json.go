package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//	1.	Обрабатываем ошибку
	if err != nil {
		errorResp := map[string]string{"error": fmt.Sprintf("ошибка обработки запроса: %v", err)}
		errorJSON, _ := json.Marshal(errorResp)
		w.Write(errorJSON)
		return
	}

	//	2.	Возвращаем ответ успешной обработки запроса
	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(data)
	w.Write(result)
}
