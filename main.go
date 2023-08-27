package main

import (
	"log"
	"os"

	BotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/markuszver/hadesBot/vars"
)

func main() {
	var bot *BotAPI.BotAPI
	//hadesChannel:=
	bot, err := BotAPI.NewBotAPI(os.Getenv("BOT_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
	chatID, err := GetChatID()
	chatCfg := BotAPI.ChatInviteLinkConfig{
		ChatConfig: BotAPI.ChatConfig{
			ChatID:             chatID,
			SuperGroupUsername: "",
		},
	}
	bot.Debug = true
	link, err := bot.GetInviteLink(chatCfg)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := BotAPI.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		chatID := update.Message.Chat.ID
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			text := update.Message.Text
			msg := BotAPI.NewMessage(chatID, "Введите пароль")
			if text == os.Getenv("TGPASSWORD") {
				//sendInvite
			} else {
				msg := BotAPI.NewMessage(chatID, vars.IncorrectPassword)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
