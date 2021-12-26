package main

import (
	"github.com/ghaliesh/slouchwatch/api"
	"github.com/ghaliesh/slouchwatch/bot"
	"github.com/ghaliesh/slouchwatch/config"
	"github.com/ghaliesh/slouchwatch/loader"
)

func main() {
	envFileLoader := loader.NewDotEnvFileLoader()
	configEnv := config.NewDotEnvConfigInstance(envFileLoader)
	tgbotApi := api.NewTgBotApiInstance()
	tgbot := bot.NewTelegramBot(configEnv, tgbotApi)

	tgbot.HandleUpdates()
}
