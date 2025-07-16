package handlers

import (
	"encoding/json"
	"my_final_project/date"
	"my_final_project/db"
	"net/http"
	"strconv"
	"strings"
)

func AddTaskHandler(store *db.SQLSchedulerStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "ожидается формат JSON"})
			return
		}

		var task db.Task

		//десериализация JSON
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "ошибка при преобразовании JSON"})
			return
		}

		//проверка дат
		if err := date.CheckDate(&task); err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "некорректная дата"})
			return
		}

		//проверка добавление задачи
		id, err := store.AddTask(&task)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "ошибка при добавлении задачи"})
			return
		}

		//fmt.Printf("DEBUG: Получена задача: %+v\n", task)

		WriteJSON(w, http.StatusOK, map[string]string{"id": strconv.FormatInt(id, 10)})

	}
}
