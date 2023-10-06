package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/entities/dto"
	vm "github.com/ESSantana/jogo-do-bicho/internal/entities/viewmodel"
)

type BetService interface {
	GetAllBets(ctx context.Context) (allBets []vm.Bet, err error)
	// Get(id string) []vm.Bet
	Create(ctx context.Context, bet dto.Bet) (vm.Bet, error)
	// Update(bet dto.Bet) vm.Bet
	// Delete(id string) bool
}
