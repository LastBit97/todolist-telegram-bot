package bot

import (
	"log"
	"strings"

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
	if err := ctx.Send("Задача добавлена"); err != nil {
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

	markup := &tele.ReplyMarkup{}
	for _, task := range todoList {

		deleteBtn := markup.Data("Удалить", "delete", task.Id.Hex())
		markup.Inline(markup.Row(deleteBtn))

		if err := ctx.Send(task.Title, markup); err != nil {
			log.Printf("error: %v", err)
			return err
		}
	}
	return nil
}

func (h *Handler) deleteTask(ctx tele.Context) error {
	btnKind := strings.TrimSpace(ctx.Args()[0])

	if btnKind == "delete" {
		id := ctx.Args()[1]
		if err := h.taskService.DeleteTask(id); err != nil {
			log.Printf("error: %v", err)
			return err
		}
		if err := ctx.Delete(); err != nil {
			log.Printf("error: %v", err)
			return err
		}
	}
	return nil
}
