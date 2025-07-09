package handlers

import (
	"my_final_project/db"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func GetTasksHandlers(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.SortTask(50)
	if err != nil {
		WriteJSON(w, map[string]string{"error": "ошибка получения задач"})
		return
	}
	WriteJSON(w, TasksResp{Tasks: tasks})

}
