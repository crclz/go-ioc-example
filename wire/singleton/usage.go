package singleton

func usage() {
	rootComponent, err := initializeRootComponent()
	if err != nil {
		panic(err)
	}

	rootComponent.chatService.Chat()
}
