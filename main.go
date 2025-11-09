package main

import (
	"log"
	g4f "reverse-engineer/G4f"
	"strings"

	"github.com/MUSTAFA-A-KHAN/telegram-bot-anime/view"
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
			go handleMessage(bot, bot2, update.Message)
		} else if update.CallbackQuery != nil {
			go handleCallbackQuery(bot, update.CallbackQuery)
		}
	}

	return nil
}
func handleMessage(bot *tgbotapi.BotAPI, bot2 *tgbotapiV5.BotAPI, message *tgbotapi.Message) {
	if strings.HasPrefix(message.Text, "/image ") {
		prompt := strings.TrimPrefix(message.Text, "/image ")
		f, res := g4f.G4f(prompt, "image")
		if !f {
			view.ReplyToMessage(bot, message.MessageID, message.Chat.ID, res)
			return
		}
		photoMsg := tgbotapiV5.NewPhoto(message.Chat.ID, tgbotapiV5.FilePath("./output/tmp.png"))
		photoMsg.ReplyToMessageID = message.MessageID
		_, err := bot2.Send(photoMsg)
		if err != nil {
			view.ReplyToMessage(bot, message.MessageID, message.Chat.ID, err.Error())

		}

	} else {
		f, res := g4f.G4f(message.Text, "text")
		if !f {
			res = "Something went wrong"
		}
		view.ReplyToMessage(bot, message.MessageID, message.Chat.ID, res)
		// view.SendMessage(bot, message.Chat.ID, res)
	}
}
func handleCallbackQuery(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
}
