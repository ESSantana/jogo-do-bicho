package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/domain/dto"
	"github.com/ESSantana/jogo-do-bicho/internal/domain/errors"
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

func (svc *BetService) Create(ctx context.Context, bet dto.Bet) (createdBet dto.Bet, err error) {
	err = bet.Validate()
	if err != nil {
		return createdBet, errors.NewValidationError(err.Error())
	}

	raffleRepo := svc.repoManager.NewRaffleRepository()
	currentRaffle, err := raffleRepo.GetCurrentEdition(ctx)
	if err != nil {
		return createdBet, errors.NewSQLError("erro ao buscar edição do sorteio vigente")
	}

	betRepo := svc.repoManager.NewBetRepository()
	var comb []string
	for _, n := range bet.BetCombination {
		comb = append(comb, fmt.Sprint(n))
	}

	now := time.Now()
	persistedID, err := betRepo.Create(ctx, entities.Bet{
		GamblerID:      bet.GamblerID,
		RaffleID:       currentRaffle.ID,
		BetModifier:    entities.BetModifier(bet.BetModifier),
		BetType:        entities.BetType(bet.BetType.Slug),
		BetPrice:       bet.BetPrice,
		BetCombination: strings.Join(comb, ","),
		CreatedAt:      &now,
	})

	if err != nil {
		return createdBet, err
	}
	bet.ID = persistedID
	bet.CreatedAt = now
	return bet, err
}

func (svc *BetService) GetAll(ctx context.Context) (allBets []dto.Bet, err error) {
	betRepo := svc.repoManager.NewBetRepository()
	items, err := betRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	allBets = make([]dto.Bet, 0)
	for _, item := range items {
		allBets = append(allBets, dto.Bet{
			ID: item.ID,
			BetType: dto.BetType{
				Slug: dto.BetTypeSlug(item.BetType),
			},
			BetPrice:       item.BetPrice,
			BetCombination: item.GetCombinationIntValues(),
			GamblerID:      item.GamblerID,
			RaffleID:       item.RaffleID,
			BetModifier:    item.BetModifier,
			CreatedAt:      *item.CreatedAt,
		})
	}
	return allBets, nil
}

func (svc *BetService) GetByID(ctx context.Context, id int64) (bet dto.Bet, err error) {
	betRepo := svc.repoManager.NewBetRepository()

	persisted, err := betRepo.GetByID(ctx, bet.ID)
	if err != nil {
		return bet, err
	}

	if !persisted.IsValid() {
		return bet, errors.NewNotFoundError("registro não encontrado")
	}

	return dto.Bet{
		ID: persisted.ID,
		BetType: dto.BetType{
			Slug: dto.BetTypeSlug(persisted.BetType),
		},
		BetPrice:       persisted.BetPrice,
		BetCombination: persisted.GetCombinationIntValues(),
		GamblerID:      persisted.GamblerID,
		RaffleID:       persisted.RaffleID,
		BetModifier:    persisted.BetModifier,
		CreatedAt:      *persisted.CreatedAt,
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
		ID:             bet.ID,
		GamblerID:      bet.GamblerID,
		RaffleID:       bet.RaffleID,
		BetModifier:    entities.BetModifier(bet.BetModifier),
		BetType:        entities.BetType(bet.BetType.Slug),
		BetPrice:       bet.BetPrice,
		BetCombination: strings.Join(comb, ","),
	})

	if err != nil {
		return false, err
	}

	if rowsAffected == 1 {
		return true, nil
	}

	return false, errors.NewSQLError("não foi possivel atualizar o registro")
}
