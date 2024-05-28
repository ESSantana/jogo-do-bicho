package services

import (
	"context"
	"errors"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/entities/dto"
	vm "github.com/ESSantana/jogo-do-bicho/internal/entities/viewmodel"
	repo_contracts "github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
	"github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
	"github.com/ESSantana/jogo-do-bicho/packages/log"
	"github.com/jackc/pgx/v5/pgtype"
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
		DocumentType: db.DocType(gambler.DocumentType),
		BirthDate:    gambler.BirthDate.Format(time.RFC3339),
	}

	gamblerRepo := svc.repoManager.NewGamblerRepository()
	persistedGambler, err := gamblerRepo.Create(ctx, gamblerParams)
	if err != nil {
		return createdGambler, err
	}

	birthDate, err := time.Parse(time.RFC3339, persistedGambler.BirthDate)
	if err != nil {
		return createdGambler, err
	}

	returnGambler := vm.Gambler{
		ID:           persistedGambler.ID,
		Name:         persistedGambler.GamblerName,
		Document:     persistedGambler.Document,
		DocumentType: string(persistedGambler.DocumentType),
		BirthDate:    birthDate,
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
		birthDate, err := time.Parse(time.RFC3339, item.Gambler.BirthDate)
		if err != nil {
			return allGamblers, err
		}

		allGamblers = append(allGamblers, vm.Gambler{
			ID:           item.Gambler.ID,
			Name:         item.Gambler.GamblerName,
			Document:     item.Gambler.GamblerName,
			DocumentType: item.Gambler.GamblerName,
			BirthDate:    birthDate,
		})
	}
	return allGamblers, nil
}

func (svc *GamblerService) GetByID(ctx context.Context, id int32) (gambler vm.Gambler, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()
	gamblerPersisted, err := gamblerRepo.GetByID(ctx, id)
	if err != nil {
		return gambler, err
	}

	birthDate, err := time.Parse(time.RFC3339, gamblerPersisted[0].Gambler.BirthDate)
	if err != nil {
		return gambler, err
	}

	gambler = vm.Gambler{
		ID:           gamblerPersisted[0].Gambler.ID,
		Name:         gamblerPersisted[0].Gambler.GamblerName,
		Document:     gamblerPersisted[0].Gambler.GamblerName,
		DocumentType: gamblerPersisted[0].Gambler.GamblerName,
		BirthDate:    birthDate,
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

	updateParams := db.UpdateGamblerParams{
		ID:           gambler.ID,
		GamblerName:  gambler.Name,
		Document:     gambler.Document,
		DocumentType: db.DocType(gambler.DocumentType),
		BirthDate:    gambler.BirthDate.Format(time.RFC3339),
	}

	updatedGambler, err := gamblerRepo.Update(ctx, updateParams)
	if err != nil {
		return false, err
	}

	if updatedGambler.ID == 0 {
		return false, errors.New("internal server error")
	}

	return true, nil
}

func (svc *GamblerService) Delete(ctx context.Context, id int32) (deleted bool, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()

	deleteParams := db.DeleteGamblerParams{
		ID: id,
		DeletedAt: pgtype.Timestamp{
			Time: time.Now(),
		},
	}

	deletedGambler, err := gamblerRepo.Delete(ctx, deleteParams)
	if err != nil {
		return false, err
	}

	if deletedGambler.ID == 0 {
		return false, errors.New("internal server error")
	}

	return true, nil
}
