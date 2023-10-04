package notification

type Notification interface {
	SendMsg(msg string) error
}

func NewNotification(accessToken string) Notification {
	return newNotify(accessToken)
}
