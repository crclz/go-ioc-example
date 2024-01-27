package scoped

import "github.com/crclz/go-ioc-example/wire/singleton"

type AuthenticationService struct {
	dbConnectionService *singleton.DbConnectionService
}

func NewAuthenticationService(
	dbConnectionService *singleton.DbConnectionService,
) *AuthenticationService {
	return &AuthenticationService{dbConnectionService: dbConnectionService}
}

func (p *AuthenticationService) Authenticate() {

}
