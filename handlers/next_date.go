package handlers

import (
	"log"
	"my_final_project/date"
	"net/http"
	"time"
)

func NextDateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "метод не поддерживается"})
		return
	}

	now := r.FormValue("now")
	startDate := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var nowTime time.Time
	var err error

	if now == "" {
		nowTime = time.Now()
	} else {
		nowTime, err = time.Parse(date.FormatDate, now)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Неверный формат даты"})
			return
		}
	}

	res, err := date.NextDate(nowTime, startDate, repeat)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(res))
	if err != nil {
		log.Printf("Ошибка записи ответа: %v", err)
	}

}
