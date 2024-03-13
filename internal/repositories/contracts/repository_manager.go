package contracts

type RepositoryManager interface {
	NewBetRepository() BetRepository
	NewGamblerRepository() GamblerRepository
}
