package scoped

import "github.com/crclz/go-ioc-example/wire/singleton"

func main() {
	var rootServiceLocator, err = singleton.InitializeRootServiceLocator()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		handleRequest(rootServiceLocator)
	}
}

func handleRequest(rootServiceLocator *singleton.RootServiceLocator) error {
	var scopedRootServiceLocator, err = initializeScopedRootServiceLocator(rootServiceLocator)
	if err != nil {
		panic(err)
	}

	scopedRootServiceLocator.authenticationService.Authenticate()
	return nil
}
