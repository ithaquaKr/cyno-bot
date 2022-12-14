package bot

import (
	"log"
	"os"
	"time"

	tb "gopkg.in/telebot.v3"
)

// Run init bot, register handlers, start bot
func Run() {

	//init bot
	pref := tb.Settings{
		Token: os.Getenv("TOKEN"),
		// Token:  "5772285972:AAGfog9OMQm1a0N0ePsEnlPKL8Rl8mRPcp4",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}
	log.Printf("start 1")

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatalf("error while creating telebot: %s", err)
		return
	}

	b.Handle("/hello", func(c tb.Context) error {
		return c.Send("Hello!")
	})
	log.Printf("start")
	//bot start loop
	b.Start()

}
