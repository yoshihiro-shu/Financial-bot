package notification_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoshihiro-shu/financial-bot/internal/notification"
)

var channelToken = os.Getenv("LINE_BOT_CHANNEL_TOKEN")
var channelSecret = os.Getenv("LINE_BOT_CHANNEL_SECRET")

func TestLineBotSendMsg(t *testing.T) {
	if channelToken == "" || channelSecret == "" {
		t.Skip()
	}
	bot := notification.NewLineBot(channelToken, channelSecret)
	err := bot.SendMsg("これはテストメッセージです。")
	if !assert.NoError(t, err) {
		t.Errorf("error is %s", err)
		t.FailNow()
	}
}
