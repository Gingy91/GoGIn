package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type TaskRequest struct {
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка JSON", http.StatusBadRequest)
		return

	}
	task := Task{Task: req.Task, IsDone: req.IsDone}
	DB.Create(&task)
	fmt.Fprintln(w, "Task обновлен")

}
func main() {
	InitDB()
	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", PostHandler).Methods("POST")
	router.HandleFunc("/api/tasks", GetHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}
