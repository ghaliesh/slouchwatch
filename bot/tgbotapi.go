package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TgBotApi interface {
	NewUpdate(offset int) tgbotapi.UpdateConfig
	NewMessage(chatID int64, text string) tgbotapi.MessageConfig
	NewBotAPI(token string) (*tgbotapi.BotAPI, error)
}

type LibTgBotApi struct {
	bot tgbotapi.BotAPI
}

func (b LibTgBotApi) NewUpdate(offset int) tgbotapi.UpdateConfig {
	return tgbotapi.NewUpdate(offset)
}

func (b LibTgBotApi) NewMessage(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, text)
}

func (b LibTgBotApi) NewBotAPI(token string) (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI(token)
}

func NewTgBotApiInstance() TgBotApi {
	return LibTgBotApi{bot: tgbotapi.BotAPI{}}
}
