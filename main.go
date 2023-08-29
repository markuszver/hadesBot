package main

import (
	"log"

	BotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/markuszver/hadesBot/config"
	"github.com/markuszver/hadesBot/handlers"
	"github.com/markuszver/hadesBot/utils"
)

func main() {
	var bot *BotAPI.BotAPI
	//hadesChannel:=
	bot, err := BotAPI.NewBotAPI(config.Config("BOT_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	link, err := utils.GetInviteLink(bot)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := BotAPI.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			chatID := update.Message.Chat.ID
			msg := BotAPI.NewMessage(chatID, "")
			// If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			switch update.Message.Command() {
			case "start":
				msg.Text = "ну привет, лобачевский сука"
			case "sendphoto":
				handlers.HandlePhoto(bot, updates)
				msg.Text = "waiting for new command..."
			case "getlink":
				//password
				msg.Text = link
			default:
				msg.Text = "wrong command"
			}
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
