package main

import (
	"github.com/roman-mazur/chat-channels-example/server/channels"
)

// Injectors from modules.go:

func ComposeApiServer(port HttpPortNumber) (*ChatApiServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	store := channels.NewStore(db)
	httpHandlerFunc := channels.HttpHandler(store)
	chatApiServer := &ChatApiServer{
		Port:            port,
		ChannelsHandler: httpHandlerFunc,
	}
	return chatApiServer, nil
}
