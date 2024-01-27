package singleton

import (
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	NewChatService,
	NewDbConnectionService,
	wire.Struct(new(RootComponent), "*"),
)

// RootComponent: put all services that you want to directly access in RootComponent
type RootComponent struct {
	chatService *ChatService
}
