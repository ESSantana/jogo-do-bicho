package services

import (
	"context"
	"sync"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/domain/dto"
	"github.com/ESSantana/jogo-do-bicho/internal/domain/errors"
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

func (svc *GamblerService) Create(ctx context.Context, gambler dto.Gambler) (createdGambler dto.Gambler, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()

	now := time.Now()
	persistedID, err := gamblerRepo.Create(ctx, entities.Gambler{
		Name:         gambler.Name,
		Document:     gambler.Document,
		DocumentType: entities.GamblersDocumentType(gambler.DocumentType),
		BirthDate:    &gambler.BirthDate,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	})

	if err != nil {
		return createdGambler, err
	}

	returnGambler := dto.Gambler{
		ID:           persistedID,
		Name:         gambler.Name,
		Document:     gambler.Document,
		DocumentType: gambler.DocumentType,
		BirthDate:    gambler.BirthDate,
	}

	return returnGambler, err
}

func (svc *GamblerService) GetAll(ctx context.Context) (allGamblers []dto.Gambler, err error) {
	gamblerRepo := svc.repoManager.NewGamblerRepository()
	gamblers, err := gamblerRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	allGamblers = make([]dto.Gambler, 0)
	for _, item := range gamblers {

		allGamblers = append(allGamblers, dto.Gambler{
			ID:           item.ID,
			Name:         item.Name,
			Document:     item.Document,
			DocumentType: string(item.DocumentType),
			BirthDate:    *item.BirthDate,
		})
	}
	return allGamblers, nil
}

func (svc *GamblerService) GetByID(ctx context.Context, id int64) (gambler dto.Gambler, err error) {
	var wg sync.WaitGroup

	gamblerChan := make(chan dto.Gambler)
	betsChan := make(chan []dto.Bet)

	getGamblerByID := func() {
		wg.Add(1)
		defer wg.Done()

		gamblerRepo := svc.repoManager.NewGamblerRepository()
		persisted, err := gamblerRepo.GetByID(ctx, id)
		if err != nil {
			return
		}
		if persisted.IsValid() {
			gamblerChan <- dto.Gambler{
				ID:           persisted.ID,
				Name:         persisted.Name,
				Document:     persisted.Document,
				DocumentType: string(persisted.DocumentType),
				BirthDate:    *persisted.BirthDate,
			}
		}
	}

	getBetByGamblerID := func() {
		wg.Add(1)
		defer wg.Done()

		betRepo := svc.repoManager.NewBetRepository()
		bets, err := betRepo.GetAllByGamblerID(ctx, id)
		if err != nil {
			return
		}

		betsVM := make([]dto.Bet, 0)
		for _, bet := range bets {
			betsVM = append(betsVM, dto.Bet{
				ID:        bet.ID,
				GamblerID: bet.GamblerID,
				RaffleID:  bet.RaffleID,
				BetType: dto.BetType{
					Slug: dto.BetTypeSlug(bet.BetType),
				},
				BetModifier:    string(bet.BetModifier),
				BetPrice:       bet.BetPrice,
				BetCombination: bet.GetCombinationIntValues(),
				CreatedAt:      *bet.CreatedAt,
			})
		}
		betsChan <- betsVM
	}

	go getGamblerByID()
	go getBetByGamblerID()

	wg.Wait()

	gambler = <-gamblerChan
	gambler.Bets = <-betsChan

	return gambler, nil
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
		Name:         gambler.Name,
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
