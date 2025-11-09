package handler

import (
	g4f "reverse-engineer/G4f"
	"strings"

	"github.com/MUSTAFA-A-KHAN/telegram-bot-anime/view"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	tgbotapiV5 "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, bot2 *tgbotapiV5.BotAPI, message *tgbotapi.Message) {
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

	} else if strings.HasPrefix(message.Text, "/say ") {
		text := strings.TrimPrefix(message.Text, "/say ")
		f, res := g4f.G4f(text, "custom-voice")
		if !f {
			view.ReplyToMessage(bot, message.MessageID, message.Chat.ID, res)
			return
		}
		audioMsg := tgbotapiV5.NewVoice(message.Chat.ID, tgbotapiV5.FilePath("./output/tmp.mp3"))
		audioMsg.ReplyToMessageID = message.MessageID
		_, err := bot2.Send(audioMsg)
		if err != nil {
			view.ReplyToMessage(bot, message.MessageID, message.Chat.ID, err.Error())

		}

	} else {
		f, res := g4f.G4f(message.Text, "text")
		if !f {
			res = "Something went wrong"
		}
		view.ReplyToMessage(bot, message.MessageID, message.Chat.ID, res)
		// view.SendMessage(bot, message.MessageID, message.Chat.ID, res)
	}
}
func HandleCallbackQuery(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
}
