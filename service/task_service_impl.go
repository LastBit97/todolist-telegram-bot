package service

import (
	"github.com/LastBit97/todolist-telegram-bot/model"
	"github.com/LastBit97/todolist-telegram-bot/repository"
)

type TaskServiceImpl struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &TaskServiceImpl{repo}
}

func (ts *TaskServiceImpl) AddTask(task *model.Task) error {
	return ts.repo.CreateTask(task)
}

func (ts *TaskServiceImpl) GetTasks(chatId int64) ([]*model.Task, error) {
	return ts.repo.GetTasks(chatId)
}

func (ts *TaskServiceImpl) DeleteTask(taskId string) error {
	return ts.repo.DeleteTask(taskId)
}
