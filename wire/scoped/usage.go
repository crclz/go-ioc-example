package scoped

import "github.com/crclz/go-ioc-example/wire/singleton"

func main() {
	var rootComponent, err = singleton.InitializeRootComponent()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		handleRequest(rootComponent)
	}
}

func handleRequest(rootComponent *singleton.RootComponent) error {
	var scopedRootComponent, err = initializeScopedRootComponent(rootComponent)
	if err != nil {
		panic(err)
	}

	scopedRootComponent.authenticationService.Authenticate()
	return nil
}
