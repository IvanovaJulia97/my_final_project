package main

import (
	"log"
	"my_final_project/db"
	rout "my_final_project/router"
	"my_final_project/server"

	_ "modernc.org/sqlite"
)

func main() {

	store, err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка в подключении к БД: %v", err)
	}

	defer store.Close()

	log.Println("Успешное подключение к БД")

	rout.Init(store)

	server.Server()
}
