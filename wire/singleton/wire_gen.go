// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package singleton

// Injectors from wire_stub.go:

func InitializeRootComponent() (*RootComponent, error) {
	dbConnectionService := NewDbConnectionService()
	chatService := NewChatService(dbConnectionService)
	rootComponent := &RootComponent{
		ChatService:         chatService,
		DbConnectionService: dbConnectionService,
	}
	return rootComponent, nil
}
