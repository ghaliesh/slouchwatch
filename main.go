package main

import (
	"github.com/ghaliesh/slouchwatch/bot"
)

func main() {
	tgbot := bot.InitBot()

	bot.HandleBotsUpdates(tgbot)
}
