package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type TaskRequest struct {
	Task string `json:"task"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello", task)

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
	InitDB()
	DB.AutoMigrate(&Task{})
	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", PostHandler).Methods("POST")
	router.HandleFunc("/api/tasks", GetHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}
