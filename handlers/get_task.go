package handlers

import (
	"my_final_project/db"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

const limit = 50

func GetTasksHandlers(store *db.SQLSchedulerStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := store.SortTask(limit)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "ошибка получения задач"})
			return
		}
		WriteJSON(w, http.StatusOK, TasksResp{Tasks: tasks})

	}
}
