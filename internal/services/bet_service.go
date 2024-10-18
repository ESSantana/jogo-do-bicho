package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/domain/dto"
	"github.com/ESSantana/jogo-do-bicho/internal/domain/errors"
	vm "github.com/ESSantana/jogo-do-bicho/internal/domain/viewmodel"
	repo_contracts "github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/entities"
	"github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
	"github.com/ESSantana/jogo-do-bicho/packages/log"
)

type BetService struct {
	logger      log.Logger
	repoManager repo_contracts.RepositoryManager
}

func newBetService(logger log.Logger, repoManager repo_contracts.RepositoryManager) contracts.BetService {
	return &BetService{
		logger:      logger,
		repoManager: repoManager,
	}
}

func (svc *BetService) Create(ctx context.Context, bet dto.Bet) (createdBet vm.Bet, err error) {
	err = bet.Validate()
	if err != nil {
		return createdBet, errors.NewValidationError(err.Error())
	}

	betRepo := svc.repoManager.NewBetRepository()

	var comb []string
	for _, n := range bet.BetCombination {
		comb = append(comb, fmt.Sprint(n))
	}

	persistedID, err := betRepo.Create(ctx, entities.Bet{
		GamblerID: bet.GamblerID,
		BetType:   string(bet.BetType.Slug),
		BetPrice:  bet.BetPrice,
		BetChoice: strings.Join(comb, ","),
	})

	if err != nil {
		return createdBet, err
	}

	returnBet := vm.Bet{
		ID:        persistedID,
		BetType:   createdBet.BetType,
		BetPrice:  createdBet.BetPrice,
		BetChoice: createdBet.BetChoice,
	}

	return returnBet, err
}

func (svc *BetService) GetAll(ctx context.Context) (allBets []vm.Bet, err error) {
	betRepo := svc.repoManager.NewBetRepository()
	items, err := betRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	allBets = make([]vm.Bet, 0)
	for _, item := range items {
		allBets = append(allBets, vm.Bet{
			ID:        item.ID,
			BetType:   item.BetType,
			BetPrice:  item.BetPrice,
			BetChoice: item.BetChoice,
			Gambler: &vm.Gambler{
				ID: item.GamblerID,
			},
		})
	}
	return allBets, nil
}

func (svc *BetService) GetByID(ctx context.Context, id int64) (bet vm.Bet, err error) {
	betRepo := svc.repoManager.NewBetRepository()

	persisted, err := betRepo.GetByID(ctx, bet.ID)
	if err != nil {
		return bet, err
	}

	if !persisted.IsValid() {
		return bet, errors.NewNotFoundError("registro não encontrado")
	}

	return vm.Bet{
		ID: persisted.ID,
		Gambler: &vm.Gambler{
			ID: persisted.GamblerID,
		},
		BetType:   persisted.BetType,
		BetPrice:  persisted.BetPrice,
		BetChoice: persisted.BetChoice,
	}, nil
}

func (svc *BetService) Update(ctx context.Context, bet dto.Bet) (updated bool, err error) {
	betRepo := svc.repoManager.NewBetRepository()

	persisted, err := betRepo.GetByID(ctx, bet.ID)
	if err != nil {
		return false, err
	}

	if !persisted.IsValid() {
		return false, errors.NewNotFoundError("registro não encontrado")
	}

	var comb []string
	for _, n := range bet.BetCombination {
		comb = append(comb, fmt.Sprint(n))
	}

	rowsAffected, err := betRepo.Update(ctx, entities.Bet{
		ID:        bet.ID,
		BetPrice:  bet.BetPrice,
		BetChoice: strings.Join(comb, ","),
		BetType:   string(bet.BetType.Slug),
	})

	if err != nil {
		return false, err
	}

	if rowsAffected == 1 {
		return true, nil
	}

	return false, errors.NewSQLError("não foi possivel atualizar o registro")
}

func (svc *BetService) Delete(ctx context.Context, id int64) (deleted bool, err error) {
	betRepo := svc.repoManager.NewBetRepository()

	now := time.Now()

	rowsAffected, err := betRepo.Delete(ctx, entities.Bet{
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
