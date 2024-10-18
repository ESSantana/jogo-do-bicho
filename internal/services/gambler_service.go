package services

import (
	"context"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/domain/dto"
	"github.com/ESSantana/jogo-do-bicho/internal/domain/errors"
	vm "github.com/ESSantana/jogo-do-bicho/internal/domain/viewmodel"
	repo_contracts "github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/entities"
	"github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
	"github.com/ESSantana/jogo-do-bicho/packages/log"
)

type GamblerService struct {
	logger      log.Logger
	repoManager repo_contracts.RepositoryManager
}

func newGamblerService(logger log.Logger, repoManager repo_contracts.RepositoryManager) contracts.GamblerService {
	return &GamblerService{
		logger:      logger,
		repoManager: repoManager,
	}
}

func (svc *GamblerService) Create(ctx context.Context, gambler dto.Gambler) (createdGambler vm.Gambler, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()

	now := time.Now()
	persistedID, err := gamblerRepo.Create(ctx, entities.Gambler{
		GamblerName:  gambler.Name,
		Document:     gambler.Document,
		DocumentType: entities.GamblersDocumentType(gambler.DocumentType),
		BirthDate:    &gambler.BirthDate,
		UpdatedAt:    &now,
	})

	if err != nil {
		return createdGambler, err
	}

	returnGambler := vm.Gambler{
		ID:           persistedID,
		Name:         gambler.Name,
		Document:     gambler.Document,
		DocumentType: gambler.DocumentType,
		BirthDate:    gambler.BirthDate,
	}

	return returnGambler, err
}

func (svc *GamblerService) GetAll(ctx context.Context) (allGamblers []vm.Gambler, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()
	gamblers, err := gamblerRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	allGamblers = make([]vm.Gambler, 0)
	for _, item := range gamblers {

		allGamblers = append(allGamblers, vm.Gambler{
			ID:           item.ID,
			Name:         item.GamblerName,
			Document:     item.Document,
			DocumentType: string(item.DocumentType),
			BirthDate:    *item.BirthDate,
		})
	}
	return allGamblers, nil
}

func (svc *GamblerService) GetByID(ctx context.Context, id int64) (gambler vm.Gambler, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()

	persisted, err := gamblerRepo.GetByID(ctx, id)
	if err != nil {
		return gambler, err
	}

	if !persisted.IsValid() {
		return gambler, errors.NewNotFoundError("registro não encontrado")
	}

	return vm.Gambler{
		ID:           persisted.ID,
		Name:         persisted.GamblerName,
		Document:     persisted.Document,
		DocumentType: string(persisted.DocumentType),
		BirthDate:    *persisted.BirthDate,
	}, nil
}

func (svc *GamblerService) Update(ctx context.Context, gambler dto.Gambler) (updated bool, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()

	persisted, err := gamblerRepo.GetByID(ctx, gambler.ID)
	if err != nil {
		return false, err
	}

	if !persisted.IsValid() {
		return false, errors.NewNotFoundError("registro não encontrado")
	}

	now := time.Now()

	rowsAffected, err := gamblerRepo.Update(ctx, entities.Gambler{
		ID:           gambler.ID,
		GamblerName:  gambler.Name,
		Document:     gambler.Document,
		DocumentType: entities.GamblersDocumentType(gambler.DocumentType),
		BirthDate:    &gambler.BirthDate,
		UpdatedAt:    &now,
	})
	if err != nil {
		return false, err
	}

	if rowsAffected == 1 {
		return true, nil
	}

	return false, errors.NewSQLError("não foi possivel atualizar o registro")
}

func (svc *GamblerService) Delete(ctx context.Context, id int64) (deleted bool, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()

	now := time.Now()

	rowsAffected, err := gamblerRepo.Delete(ctx, entities.Gambler{
		ID:        id,
		DeletedAt: &now,
	})

	if err != nil {
		return false, err
	}

	if rowsAffected == 1 {
		return true, nil
	}

	return false, errors.NewSQLError("não foi possivel deletar o registro")
}
