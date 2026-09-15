package main

import (
	"fmt"
	"net/http"

	"github.com/Bal1sta/TaskFlow/handler"
	"github.com/Bal1sta/TaskFlow/model"
)

func main() {

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/tasks", handler.TasksHandler)
	http.HandleFunc("/health", handler.HealthHandler)

	task := model.Task{
		ID:     1,
		Title:  "Изучить GO",
		Status: "new",
	}
	fmt.Println(task)

	fmt.Println("Сервер запущен!")
	http.ListenAndServe(":8080", nil)
}
