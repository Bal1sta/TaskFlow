package model

import (
	"fmt"
)

type Task struct {
	ID     int
	Title  string
	Status string
}

func (t Task) Print() {
	fmt.Println("Задача:", t.Title)
	fmt.Println("Статус:", t.Status)
	fmt.Println("ID:", t.ID)
}