package bot

import (
	"log"
	"strconv"

	"github.com/ghaliesh/slouchwatch/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot interface {
	Init(envloader config.EnvLoader)
	HandleUpdates()
}

type TelegarmBot struct {
	bot *tgbotapi.BotAPI
}

func (b *TelegarmBot) Init(envloader config.EnvLoader, tgbotapi TgBotApi) {
	token, _ := config.GetEnvVar(envloader, "BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	debugstr, _ := config.GetEnvVar(envloader, "DEBUG")
	debug, error := strconv.ParseBool(debugstr)
	if error != nil {
		log.Fatalf("DEBUG env variable is either missing or assigned to non-bool value, more details below:\n %v", error.Error())
	}
	bot.Debug = debug

	b.bot = bot
}

func (b *TelegarmBot) HandleUpdates(tgbotapi TgBotApi) {

	bot := b.bot
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf(update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func NewTelegramBot() *TelegarmBot {
	return &TelegarmBot{}
}
