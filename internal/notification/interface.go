package notification

type Notification interface {
	SendMsg(msg string) error
}