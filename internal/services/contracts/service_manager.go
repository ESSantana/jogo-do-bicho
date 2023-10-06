package contracts

type ServiceManager interface {
	NewBetService() BetService
}
