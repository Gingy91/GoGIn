package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var task string

type TaskRequest struct {
	Task string `json:"task"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Hello", task)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод GET")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req TaskRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Ошибка JSON", http.StatusBadRequest)
			return

		}
		task = req.Task
		fmt.Fprintln(w, "Task обновлен")
	} else {
		http.Error(w, "Поддерживается только метод POST", http.StatusBadRequest)
	}
}
func main() {
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/post", PostHandler)

	http.ListenAndServe("localhost:8080", nil)

}
