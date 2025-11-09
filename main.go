package main

import (
	"log"
	"reverse-engineer/handler"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	tgbotapiV5 "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	StartBot("token")
}

// StartBot initializes and starts the bot
func StartBot(token string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	bot2, err := tgbotapiV5.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	// Pass the client instance to handleMessage and handleCallbackQuery via closure
	for update := range updates {
		if update.Message != nil {
			go handler.HandleMessage(bot, bot2, update.Message)
		} else if update.CallbackQuery != nil {
			go handler.HandleCallbackQuery(bot, update.CallbackQuery)
		}
	}

	return nil
}
