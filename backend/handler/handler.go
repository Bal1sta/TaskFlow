package handler

import (
	"fmt"
	"net/http"

	"github.com/Bal1sta/TaskFlow/model"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TaskFlow API. Главная страница")
}
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	task := model.Task{
		ID: 1,
		Title: "Изучить GO",
		Status: "new",
	}
	fmt.Fprintln(w, task.Title)
	fmt.Fprintln(w, task.Status)
	fmt.Fprintln(w, task.ID)
}
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TaskFlow API. Проверка работы сервера")
}


