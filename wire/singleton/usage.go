package singleton

func usage() {
	rootComponent, err := InitializeRootComponent()
	if err != nil {
		panic(err)
	}

	rootComponent.ChatService.Chat()
}
