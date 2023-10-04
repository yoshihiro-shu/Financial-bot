package notification

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// notification using by Line Notify
// https://notify-bot.line.me/en/
type notify struct {
	accessToken string
}

const notifyURL = "https://notify-api.line.me/api/notify"

func newNotify(accessToken string) *notify {
	return &notify{
		accessToken: accessToken,
	}
}

func (n *notify) SendMsg(msg string) error {
	u, err := url.ParseRequestURI(notifyURL)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("message", msg)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest(http.MethodPost, u.String(), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.accessToken))

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		if res.StatusCode > 299 {
			return fmt.Errorf("status code: %d", res.StatusCode)
		}
		return err
	}

	return nil
}
