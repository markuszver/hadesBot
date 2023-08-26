package main

import (
BotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
"log"
"github.com/markuszver/hadesBot/vars"
"os"
)

func main() {
	
	//hadesChannel:= 
	bot, err := BotAPI.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}
	// chatCfg:= BotAPI.ChatInviteLinkConfig{
	// 	ChatConfig: BotAPI.ChatConfig{
	// 		ChatID: 1,
	// 		SuperGroupUsername: "123",
	// 	},
	// }
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := BotAPI.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			text:= update.Message.Text
			if text == os.Getenv("TGPASSWORD") {
				//sendInvite
			} else {
				msg := BotAPI.NewMessage(update.Message.Chat.ID, vars.IncorrectPassword)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
}
}