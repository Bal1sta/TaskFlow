package main

import (
	"fmt"
	"net/http"
)



func ServerHTTP(w http.ResponseWriter, r *http.Request){

	if r.URL.Path == "/tasks" {
		fmt.Fprintln(w, "TaskFlow API. Список задач")
	} else {
		fmt.Fprintln(w, "TaskFlow API. Главная страница")
	}
}

func main() {


	http.HandleFunc("/", ServerHTTP)

  fmt.Println("Сервер запущен!")
  http.ListenAndServe(":8080", nil)
}