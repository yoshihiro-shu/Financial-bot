package notification

func NewNotification(accessToken string) Notification {
	return newNotify(accessToken)
}
