package main

import (
	"log"

	"github.com/ithaquaKr/cyno-bot/pkg/bot"
)

func main() {

	go bot.Run()
	log.Printf("Bot started")

	// go http.RunServer()
	// log.Printf("Http server started")

	select {}
}
