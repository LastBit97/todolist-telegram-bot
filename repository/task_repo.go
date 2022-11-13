package repository

import "github.com/LastBit97/todolist-telegram-bot/model"

type TaskRepository interface {
	CreateTask(task *model.Task) error
	GetTasks(chatId string) ([]*model.Task, error)
}
