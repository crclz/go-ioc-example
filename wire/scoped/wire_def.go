package scoped

import (
	"github.com/crclz/go-ioc-example/wire/singleton"
	"github.com/google/wire"
)

var ScopedSet = wire.NewSet(
	wire.Struct(new(ScopedRootDependencyLocator), "*"),

	NewAuthenticationService,

	// from root component provide ...
	FromRootDependencyLocatorProvideDbConnectionService,
)

// ScopedRootDependencyLocator: put all services that you want to directly access in here
type ScopedRootDependencyLocator struct {
	authenticationService *AuthenticationService
}

func FromRootDependencyLocatorProvideDbConnectionService(p *singleton.RootDependencyLocator) *singleton.DbConnectionService {
	return p.DbConnectionService
}
