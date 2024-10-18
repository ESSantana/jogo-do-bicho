package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/domain/dto"
	"github.com/ESSantana/jogo-do-bicho/internal/domain/errors"
	vm "github.com/ESSantana/jogo-do-bicho/internal/domain/viewmodel"
	repo_contracts "github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
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
	gamblerParams := db.CreateGamblerParams{
		GamblerName:  gambler.Name,
		Document:     gambler.Document,
		DocumentType: db.GamblersDocumentType(gambler.DocumentType),
		BirthDate:    gambler.BirthDate,
	}

	gamblerRepo := svc.repoManager.NewGamblerRepository()
	err = gamblerRepo.Create(ctx, gamblerParams)
	if err != nil {
		return createdGambler, err
	}

	//TODO: Retrieve ID
	returnGambler := vm.Gambler{
		ID:           19999999,
		Name:         gambler.Name,
		Document:     gambler.Document,
		DocumentType: gambler.DocumentType,
		BirthDate:    gambler.BirthDate,
	}

	return returnGambler, err
}

func (svc *GamblerService) GetAll(ctx context.Context) (allGamblers []vm.Gambler, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()
	items, err := gamblerRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	allGamblers = make([]vm.Gambler, 0)
	for _, item := range items {

		allGamblers = append(allGamblers, vm.Gambler{
			ID:           item.Gambler.ID,
			Name:         item.Gambler.GamblerName,
			Document:     item.Gambler.Document,
			DocumentType: string(item.Gambler.DocumentType),
			BirthDate:    item.Gambler.BirthDate,
		})
	}
	return allGamblers, nil
}

func (svc *GamblerService) GetByID(ctx context.Context, id int64) (gambler vm.Gambler, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()
	gamblerPersisted, err := gamblerRepo.GetByID(ctx, id)
	if err != nil {
		return gambler, err
	}

	gambler = vm.Gambler{
		ID:           gamblerPersisted[0].Gambler.ID,
		Name:         gamblerPersisted[0].Gambler.GamblerName,
		Document:     gamblerPersisted[0].Gambler.Document,
		DocumentType: string(gamblerPersisted[0].Gambler.DocumentType),
		BirthDate:    gamblerPersisted[0].Gambler.BirthDate,
		Bets:         []vm.Bet{},
	}

	for _, item := range gamblerPersisted {
		gambler.Bets = append(gambler.Bets, vm.Bet{
			ID:        item.Bet.ID,
			BetType:   item.Bet.BetType,
			BetPrice:  item.Bet.BetPrice,
			BetChoice: item.Bet.BetChoice,
		})
	}

	return gambler, nil
}

func (svc *GamblerService) Update(ctx context.Context, gambler dto.Gambler) (updated bool, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()

	persisted, err := gamblerRepo.GetByID(ctx, gambler.ID)
	if err != nil {
		return false, err
	}

	if len(persisted) == 0 {
		return false, errors.NewNotFoundError("apostador não encontrado")
	}

	updateParams := db.UpdateGamblerParams{
		ID:           gambler.ID,
		GamblerName:  gambler.Name,
		Document:     gambler.Document,
		DocumentType: db.GamblersDocumentType(gambler.DocumentType),
		BirthDate:    &gambler.BirthDate,
	}

	err = gamblerRepo.Update(ctx, updateParams)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (svc *GamblerService) Delete(ctx context.Context, id int64) (deleted bool, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()

	persisted, err := gamblerRepo.GetByID(ctx, id)
	if err != nil {
		return false, err
	}

	if len(persisted) == 0 {
		return false, errors.NewNotFoundError("apostador não encontrado")
	}

	deleteParams := db.DeleteGamblerParams{
		ID: id,
		DeletedAt: sql.NullTime{
			Time: time.Now(),
		},
	}

	err = gamblerRepo.Delete(ctx, deleteParams)
	if err != nil {
		return false, err
	}

	return true, nil
}
