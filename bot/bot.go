package bot

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func StartBot(apiToken string) {
	pref := tele.Settings{
		Token:  apiToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Привет")
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		return c.Reply(c.Text())
	})

	log.Print("listen to telegram api")

	b.Start()
}
