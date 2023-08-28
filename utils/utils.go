package utils

import (
	"strconv"

	BotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/markuszver/hadesBot/config"
)

type Bot struct {
	*BotAPI.BotAPI
}

func getChatID() (int64, error) {
	// url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", config.Config("BOT_APITOKEN"))
	// req, err := http.NewRequest(http.MethodPost, url, nil)
	// if err != nil {
	// 	return 0, err
	// }
	// client := &http.Client{}
	// res, err := client.Do(req)
	// if err != nil {
	// 	return 0, err
	// }
	// defer res.Body.Close()
	// b, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	return 0, err
	// }
	// result := gjson.Get(string(b), "result.#(my_chat_member.chat.title==hades).my_chat_member.chat.id").Int()
	// log.Printf("chatID: %v", result)

	// return result, nil
	i, err := strconv.ParseInt(config.Config("CHAT_ID"), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func GetInviteLink(bot *BotAPI.BotAPI) (string, error) {
	chatID, err := getChatID()
	if err != nil {
		return "", err
	}
	chatCfg := BotAPI.ChatInviteLinkConfig{
		ChatConfig: BotAPI.ChatConfig{
			ChatID:             chatID,
			SuperGroupUsername: "",
		},
	}

	return bot.GetInviteLink(chatCfg)
}
