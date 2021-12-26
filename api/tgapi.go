package api

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type LibTgBotApi struct {
	bot tgbotapi.BotAPI
}

func (b *LibTgBotApi) NewUpdate(offset int) tgbotapi.UpdateConfig {
	return tgbotapi.NewUpdate(offset)
}

func (b *LibTgBotApi) NewMessage(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, text)
}

func (b *LibTgBotApi) NewBotAPI(token string) (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI(token)
}

func NewTgBotApiInstance() BotApi {
	return &LibTgBotApi{bot: tgbotapi.BotAPI{}}
}
