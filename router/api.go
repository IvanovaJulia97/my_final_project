package rout

import (
	"my_final_project/db"
	"my_final_project/handlers"
	"net/http"
)

func Init(store *db.SQLSchedulerStore) {
	http.HandleFunc("/api/nextdate", handlers.NextDateHandler)
	http.HandleFunc("/api/task", handlers.TaskHandler(store))
	http.HandleFunc("/api/tasks", handlers.GetTasksHandlers(store))
	http.HandleFunc("/api/task/done", handlers.DoneTaskHandler(store))

}
