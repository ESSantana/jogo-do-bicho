package services

import (
	repo_contracts "github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
)

type ServiceManager struct {
	repoManager repo_contracts.RepositoryManager
}

func NewServiceManager(repoManager repo_contracts.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		repoManager: repoManager,
	}
}

func (svc *ServiceManager) NewBetService() contracts.BetService {
	return newBetService(svc.repoManager)
}
