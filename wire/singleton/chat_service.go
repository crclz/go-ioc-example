package singleton

type ChatService struct {
	dbConnectionService *DbConnectionService
}

func NewChatService(
	dbConnectionService *DbConnectionService,
) *ChatService {
	return &ChatService{
		dbConnectionService: dbConnectionService,
	}
}

func (p *ChatService) Chat() {
	p.dbConnectionService.Connect()

	p.dbConnectionService.Write("Hello")
}
