package scoped

import (
	"github.com/crclz/go-ioc-example/wire/singleton"
	"github.com/google/wire"
)

var ScopedSet = wire.NewSet(
	wire.Struct(new(ScopedRootServiceLocator), "*"),

	NewAuthenticationService,

	// from root component provide ...
	FromRootServiceLocatorProvideDbConnectionService,
)

// ScopedRootServiceLocator: put all services that you want to directly access in here
type ScopedRootServiceLocator struct {
	authenticationService *AuthenticationService
}

func FromRootServiceLocatorProvideDbConnectionService(p *singleton.RootServiceLocator) *singleton.DbConnectionService {
	return p.DbConnectionService
}
