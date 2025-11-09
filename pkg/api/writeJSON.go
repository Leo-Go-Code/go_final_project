package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, data interface{}, err error) {
	//	0.	Обрабатываем ошибку
	if err != nil {
		errorResp := map[string]string{"error": fmt.Sprintf("ошибка обработки запроса: %v", err)}
		errorJSON, _ := json.Marshal(errorResp)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorJSON)
		return
	}

	//	1.	Сериализация данных в JSON
	result, err := json.Marshal(data)
	if err != nil {
		errorResp := map[string]string{"error": "ошибка сериализации данных, полученных из БД"}
		errorJSON, _ := json.Marshal(errorResp)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorJSON)
		return
	}

	//	2.	Проставляем заголовки
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//	3.	Возвращаем ответ
	w.Write(result)
}
