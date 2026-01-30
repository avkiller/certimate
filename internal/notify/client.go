package notify

import (
	"log/slog"
)

type Client struct {
	logger *slog.Logger
}

type ClientConfigure func(*Client)

func NewClient(configures ...ClientConfigure) *Client {
	client := &Client{}
	for _, configure := range configures {
		configure(client)
	}
	return client
}

func WithLogger(logger *slog.Logger) ClientConfigure {
	return func(c *Client) {
		c.logger = logger
	}
}
