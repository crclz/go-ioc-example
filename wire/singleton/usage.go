package singleton

func usage() {
	rootDependencyLocator, err := InitializeRootDependencyLocator()
	if err != nil {
		panic(err)
	}

	rootDependencyLocator.ChatService.Chat()
}
