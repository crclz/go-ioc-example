package singleton

import (
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	NewChatService,
	NewDbConnectionService,
	wire.Struct(new(RootDependencyLocator), "*"),
)

// RootDependencyLocator: put all services that you want to directly access
type RootDependencyLocator struct {
	ChatService         *ChatService
	DbConnectionService *DbConnectionService
}
