package handlers

import (
	"my_final_project/date"
	"my_final_project/db"
	"net/http"
	"time"
)

func DoneTaskHandler(store *db.SQLSchedulerStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		if id == "" {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "не указан id задачи"})
			return
		}

		task, err := store.GetTasks(id)
		if err != nil {
			WriteJSON(w, http.StatusNotFound, map[string]string{"error": "задачи с таким id не найдено"})
			return
		}

		if task.Repeat == "" {
			if err := store.DeleteTask(id); err != nil {
				WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "ошибка правила повтора"})
				return
			}
			WriteJSON(w, http.StatusOK, map[string]string{})
			return
		}

		nextDate, err := date.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "некорректная дата"})
			return
		}

		if err := store.UpdateDate(nextDate, id); err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "ошибка обновления даты"})
			return
		}

		WriteJSON(w, http.StatusOK, map[string]string{})

	}
}
