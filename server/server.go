package server

import (
	"log"
	"net/http"
	"os"
)

const myPort = "7540"
const webDir = "./web"

func Server() {

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = myPort
	}

	file := http.FileServer(http.Dir(webDir))
	http.Handle("/", file)

	log.Printf("Сервер запущен, порт %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}

}
