package main

import (
	"fmt"
	"net/http"
)



func ServerHTTP(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, "TaskFlow API работает!")
	fmt.Fprintln(w, "Метод:", r.Method)
	fmt.Fprintln(w, "Путь:", r.URL.Path)
}

func main() {


	http.HandleFunc("/", ServerHTTP)

  fmt.Println("Сервер запущен!")
  http.ListenAndServe(":8080", nil)
}