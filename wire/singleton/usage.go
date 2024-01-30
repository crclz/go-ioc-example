package singleton

func usage() {
	rootServiceLocator, err := InitializeRootServiceLocator()
	if err != nil {
		panic(err)
	}

	rootServiceLocator.ChatService.Chat()
}
