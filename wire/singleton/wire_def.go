package singleton

import (
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	NewChatService,
	NewDbConnectionService,
	wire.Struct(new(RootServiceLocator), "*"),
)

// RootServiceLocator: put all services that you want to directly access
type RootServiceLocator struct {
	ChatService         *ChatService
	DbConnectionService *DbConnectionService
}
