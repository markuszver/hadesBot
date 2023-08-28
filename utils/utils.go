package utils

import (
	"fmt"
	"io"
	"net/http"

	BotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/markuszver/hadesBot/config"
	"github.com/tidwall/gjson"
)

type Bot struct {
	BotAPI.BotAPI
}

func GetChatID() (int64, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", config.Config("BOT_APITOKEN"))
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return 0, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	return gjson.Get(string(b), "result.#(my_chat_member.chat.title==hades).my_chat_member.chat.id").Int(), nil
}

func (bot *Bot) GetInviteLink(chatID int64) (string, error) {
	chatCfg := BotAPI.ChatInviteLinkConfig{
		ChatConfig: BotAPI.ChatConfig{
			ChatID:             chatID,
			SuperGroupUsername: "",
		},
	}

	return bot.BotAPI.GetInviteLink(chatCfg)
}
