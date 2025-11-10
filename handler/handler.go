package handler

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	g4f "reverse-engineer/G4f"
	"strings"

	"github.com/MUSTAFA-A-KHAN/telegram-bot-anime/view"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	tgbotapiV5 "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, bot2 *tgbotapiV5.BotAPI, message *tgbotapi.Message) {
	fmt.Println("Message object", message)
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

	} else if strings.HasPrefix(message.Text, "/say") {
		var text string
		if message.ReplyToMessage != nil {
			text = message.ReplyToMessage.Text
			message.MessageID = message.ReplyToMessage.MessageID
		} else {
			text = strings.TrimPrefix(message.Text, "/say ")
		}
		go bot2.Send(tgbotapiV5.NewChatAction(message.Chat.ID, tgbotapiV5.ChatRecordVoice))

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

	} else if message.Voice != nil {
		fmt.Println("Audio message received")
		//get the file contents of the audio file and send it to the model
		fileID := message.Voice.FileID

		fileURL, err := bot.GetFileDirectURL(fileID)
		if err != nil {
			fmt.Println("Error getting file URL:", err)
			return
		}
		resp, err := http.Get(fileURL)
		if err != nil {
			fmt.Println("Error downloading file:", err)
			return
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		encoded := base64.StdEncoding.EncodeToString(data)
		prompt := encoded

		fmt.Println("Got this data:", prompt)

		f, _ := g4f.G4f(string(prompt), "voice")
		if !f {
			view.ReplyToMessage(bot, message.MessageID, message.Chat.ID, "something went wrong")
			return
		}
		audioMsg := tgbotapiV5.NewVoice(message.Chat.ID, tgbotapiV5.FilePath("./output/tmp.mp3"))
		audioMsg.ReplyToMessageID = message.MessageID
		_, err = bot2.Send(audioMsg)
		if err != nil {
			view.ReplyToMessage(bot, message.MessageID, message.Chat.ID, err.Error())

		}

	} else {
		f, res := g4f.G4f(message.Text, "text")
		if !f {
			res = "Something went wrong"
		}
		view.ReplyToMessage(bot, message.MessageID, message.Chat.ID, res)
	}
}

func HandleCallbackQuery(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
}

func HandleInlineQuery(bot *tgbotapi.BotAPI, inlineQuery *tgbotapi.InlineQuery) {
	if strings.Contains(inlineQuery.Query, "/g4f") {
		fmt.Println("inside inline Query")
		prompt := strings.Trim(inlineQuery.Query, "/g4f ")
		fmt.Println(prompt)

		f, res := g4f.G4f(prompt, "text")
		if !f {
			res = "Something went wrong"
		}
		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: inlineQuery.ID,
			Results:       []interface{}{tgbotapi.NewInlineQueryResultArticleMarkdown(inlineQuery.ID, res, res)},
			CacheTime:     0,
			IsPersonal:    true,
		}
		_, err := bot.AnswerInlineQuery(inlineConf)
		if err != nil {
			fmt.Println("Error answering inline query:", err)
		}
	}

}
