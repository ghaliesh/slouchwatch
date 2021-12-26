package api

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BotApi interface {
	NewUpdate(offset int) tgbotapi.UpdateConfig
	NewMessage(chatID int64, text string) tgbotapi.MessageConfig
	NewBotAPI(token string) (*tgbotapi.BotAPI, error)
}
