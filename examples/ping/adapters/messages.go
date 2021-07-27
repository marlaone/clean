package adapters

import "github.com/marlaone/clean/interfaces"

type MessagesAdapter struct {
	messages map[string]string
}

var _ interfaces.StorageAdapter = &MessagesAdapter{}

func NewMessagesAdapter() *MessagesAdapter {
	return &MessagesAdapter{
		messages: map[string]string{
			"pong": "pong",
		},
	}
}

func (a *MessagesAdapter) Query(name string) string {
	return a.messages[name]
}
