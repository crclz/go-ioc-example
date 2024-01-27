package singleton

type DbConnectionService struct {
}

func NewDbConnectionService() *DbConnectionService {
	return &DbConnectionService{}
}

func (p *DbConnectionService) Connect() {
}

func (p *DbConnectionService) Write(s string) {
}
