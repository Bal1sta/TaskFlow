package handler

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TaskFlow API. Главная страница")
}
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TaskFlow API. Список задач")
}
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TaskFlow API. Проверка работы сервера")
}
