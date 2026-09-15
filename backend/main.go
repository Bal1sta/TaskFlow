package main

import (
	"fmt"
	"net/http"
)



func ServerHTTP(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "TaskFlow API работает!")
}

func main() {


	http.HandleFunc("/", ServerHTTP)

  fmt.Println("Сервер запущен!")
  http.ListenAndServe(":8080", nil)
}