package repository

import (

	"github.com/Bal1sta/TaskFlow/model"
)

type TaskRepository struct {
	tasks []model.Task
}
// создаёт пустой Repository
func NewTaskRepository() TaskRepository {
	return TaskRepository{tasks: []model.Task{}}
}
// добавляет задачу
func (r *TaskRepository) Add(task model.Task) {
	r.tasks = append(r.tasks, task)
}
// возвращает все задачи
func (r *TaskRepository) GetAll() []model.Task {
	return r.tasks
}