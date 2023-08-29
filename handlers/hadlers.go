package handlers

import (
	BotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/markuszver/hadesBot/utils"
)

func HandlePhoto(bot *BotAPI.BotAPI, updates BotAPI.UpdatesChannel) error {
	for update := range updates {
		if update.Message.Command() == "exit" {
			return nil
		}
		chatID := update.Message.Chat.ID
		msgID := update.Message.MessageID
		sendTo, err := utils.GetChatID()
		if err != nil {
			return err
		}
		bot.CopyMessage(BotAPI.NewCopyMessage(sendTo, chatID, msgID))
	}
	return nil
}
