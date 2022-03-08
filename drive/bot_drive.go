package drive

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type BotDb struct {
	BotTel *tgbotapi.BotAPI
}

var Bot = &BotDb{}

func ConnectBot(token string) *BotDb {
	bot, err := tgbotapi.NewBotAPI("5204140121:AAFky6KMUqdAUhvWVUPBWoOghqH4cH8lW4c")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	Bot.BotTel = bot
	return Bot
}
