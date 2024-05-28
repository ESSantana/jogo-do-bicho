package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/entities/dto"
	vm "github.com/ESSantana/jogo-do-bicho/internal/entities/viewmodel"
)

type BetService interface {
	GetAll(ctx context.Context) (bets []vm.Bet, err error)
	GetByID(ctx context.Context, id int32) (bet vm.Bet, err error)
	Create(ctx context.Context, bet dto.Bet) (betVm vm.Bet, err error)
	Update(ctx context.Context, bet dto.Bet) (updated bool, err error)
	Delete(ctx context.Context, id int32) (deleted bool, err error)
}
