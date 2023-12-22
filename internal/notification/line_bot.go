package notification

import "github.com/line/line-bot-sdk-go/v8/linebot"

type lineBot struct {
	channelToken  string
	channelSecret string
}

func NewLineBot(channelToken, channelSecret string) *lineBot {
	return &lineBot{
		channelToken:  channelToken,
		channelSecret: channelSecret,
	}
}

func (l *lineBot) SendMsg(msg string) error {
	bot, err := linebot.New(l.channelSecret, l.channelToken)
	if err != nil {
		return err
	}
	messagee := linebot.NewTextMessage(msg)
	_, err = bot.BroadcastMessage(messagee).Do()
	if err != nil {
		return err
	}
	return nil
}
