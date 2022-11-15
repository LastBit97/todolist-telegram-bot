package bot

import (
	"log"

	"github.com/LastBit97/todolist-telegram-bot/model"
	"github.com/LastBit97/todolist-telegram-bot/service"
	tele "gopkg.in/telebot.v3"
)

type Handler struct {
	taskService service.TaskService
}

func NewHandler(service service.TaskService) Handler {
	return Handler{service}
}

func (h *Handler) addTask(ctx tele.Context) error {
	task := &model.Task{
		Title:  ctx.Text(),
		ChatId: ctx.Chat().ID,
	}

	if err := h.taskService.AddTask(task); err != nil {
		log.Printf("error: %v", err)
		return err
	}
	if err := ctx.Send("Task added"); err != nil {
		log.Printf("error: %v", err)
		return err
	}
	return nil
}

func (h *Handler) getTodoList(ctx tele.Context) error {
	chatId := ctx.Chat().ID
	todoList, err := h.taskService.GetTasks(chatId)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	for _, task := range todoList {
		if err := ctx.Send(task.Title); err != nil {
			log.Printf("error: %v", err)
			return err
		}
	}
	return nil
}
