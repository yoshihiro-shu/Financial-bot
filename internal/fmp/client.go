package fmp

type client struct {
	apiKey string
}

func NewClient(apiKey string) *client {
	return &client{
		apiKey: apiKey,
	}
}
