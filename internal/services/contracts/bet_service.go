package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/entities/dto"
	vm "github.com/ESSantana/jogo-do-bicho/internal/entities/viewmodel"
)

type BetService interface {
	GetAllBets(ctx context.Context) (bets []vm.Bet, err error)
	GetByID(ctx context.Context, id int32) (bet vm.Bet, err error)
	Create(ctx context.Context, bet dto.Bet) (betVm vm.Bet, err error)
	// Update(bet dto.Bet) (betVM vm.Bet, err error)
	// Delete(id string) (deleted bool, err error)
}
