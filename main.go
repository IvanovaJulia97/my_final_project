package main

import (
	"log"
	"my_final_project/db"
	"my_final_project/rout"
	"my_final_project/server"

	_ "modernc.org/sqlite"
)

func main() {
	// err1 := godotenv.Load()
	// if err1 != nil {
	// 	log.Println("Ошибка загрузки .env файда")
	// }

	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка в подключении к БД: %v", err)
	}

	log.Println("Успешное подключение к БД")

	rout.Init()

	server.MyServer()
}
