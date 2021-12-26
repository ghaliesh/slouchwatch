package bot

import (
	"log"
	"strconv"

	"github.com/ghaliesh/slouchwatch/api"
	"github.com/ghaliesh/slouchwatch/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var configLoader config.EnvConfig
var tgapi api.BotApi

type TelegarmBot struct {
	bot *tgbotapi.BotAPI
}

func (b *TelegarmBot) HandleUpdates() {

	bot := b.bot
	u := tgapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf(update.Message.Text)

			msg := tgapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func NewTelegramBot(c config.EnvConfig, tg api.BotApi) Bot {
	// Initialize depandancies
	tgapi = tg
	configLoader = c

	token, botTokenError := configLoader.GetEnvVar("BOT_TOKEN")
	if botTokenError != nil {
		log.Fatalf("BOT_TOKEN env variable is missing :\n %v", botTokenError.Error())
	}

	bot, botInitError := tgapi.NewBotAPI(token)
	if botInitError != nil {
		log.Fatalf("Error initilizing the bot :\n %v", botTokenError.Error())
	}

	debugstr, botDebugError := configLoader.GetEnvVar("DEBUG")
	if botDebugError != nil {
		log.Fatalf("DEBUG env variable failed to be loaded :\n %v", botTokenError.Error())
	}

	debug, parseError := strconv.ParseBool(debugstr)
	if parseError != nil {
		log.Fatalf("DEBUG env variable is either missing or assigned to non-bool value, more details below:\n %v", parseError.Error())
	}

	bot.Debug = debug
	return &TelegarmBot{bot: bot}
}
