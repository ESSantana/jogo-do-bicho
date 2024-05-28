package services

import (
	repo_contracts "github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
	"github.com/ESSantana/jogo-do-bicho/packages/log"
)

type ServiceManager struct {
	logger      log.Logger
	repoManager repo_contracts.RepositoryManager
}

func NewServiceManager(logger log.Logger, repoManager repo_contracts.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		logger:      logger,
		repoManager: repoManager,
	}
}

func (svc *ServiceManager) NewBetService() contracts.BetService {
	return newBetService(svc.logger, svc.repoManager)
}

func (svc *ServiceManager) NewGamblerService() contracts.GamblerService {
	return newGamblerService(svc.logger, svc.repoManager)
}
