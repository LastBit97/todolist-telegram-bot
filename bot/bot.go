package bot

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func (h *Handler) StartBot(apiToken string) {
	pref := tele.Settings{
		Token:  apiToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	btnTasks := menu.Text("ℹ Список задач")

	menu.Reply(menu.Row(btnTasks))

	bot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send("Hello", menu)
	})

	bot.Handle(tele.OnText, h.addTask)

	bot.Handle(&btnTasks, h.getTodoList)

	log.Print("listen to telegram api")
	bot.Start()
}
