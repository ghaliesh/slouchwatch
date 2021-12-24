package main

import (
	"github.com/ghaliesh/slouchwatch/bot"
	"github.com/ghaliesh/slouchwatch/config"
)

func main() {
	tgbot := bot.NewTelegramBot()
	tgbotApi := bot.NewTgBotApiInstance()
	envloader := config.NewEnvLoaderInstance()

	tgbot.Init(envloader, tgbotApi)
	tgbot.HandleUpdates(tgbotApi)
}
