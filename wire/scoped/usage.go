package scoped

// func main() {
// 	var RootDependencyLocator, err = singleton.InitializeRootDependencyLocator()
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i := 0; i < 10; i++ {
// 		handleRequest(RootDependencyLocator)
// 	}
// }

// func handleRequest(RootDependencyLocator *singleton.RootDependencyLocator) error {
// 	var scopedRootDependencyLocator, err = initializeScopedRootDependencyLocator(RootDependencyLocator)
// 	if err != nil {
// 		panic(err)
// 	}

// 	scopedRootDependencyLocator.authenticationService.Authenticate()
// 	return nil
// }
