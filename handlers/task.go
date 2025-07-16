package handlers

import (
	"encoding/json"
	"log"
	"my_final_project/date"
	"my_final_project/db"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("ошибка кодирования JSON: %v", err)
	}
}

func TaskHandler(store *db.SQLSchedulerStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			AddTaskHandler(store)(w, r)
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			if id == "" {
				WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "не передан id задачи"})
				return
			}
			task, err := store.GetTasks(id)
			if err != nil {
				WriteJSON(w, http.StatusNotFound, map[string]string{"error": "ошибка получения задачи"})
				return
			}
			WriteJSON(w, http.StatusOK, task)

		case http.MethodPut:
			var task db.Task
			if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
				WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "ошибка преобразования в JSON"})
				return
			}
			if err := date.CheckDate(&task); err != nil {
				WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "некорректная дата"})
				return
			}
			err := store.UpdateTask(&task)
			if err != nil {
				WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "ошибка при обновлении задачи"})
				return
			}
			WriteJSON(w, http.StatusOK, map[string]string{})
		case http.MethodDelete:
			id := r.URL.Query().Get("id")
			if id == "" {
				WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "не указан id задачи"})
				return
			}
			err := store.DeleteTask(id)
			if err != nil {
				WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "ошибка при удалении задачи"})
				return
			}
			WriteJSON(w, http.StatusOK, map[string]string{})

		default:
			WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "данный запрос не поддерживается"})
		}
	}
}
