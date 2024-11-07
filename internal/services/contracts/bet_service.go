package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/domain/dto"
)

type BetService interface {
	GetAll(ctx context.Context) (bets []dto.Bet, err error)
	GetByID(ctx context.Context, id int64) (bet dto.Bet, err error)
	Create(ctx context.Context, bet dto.Bet) (betOut dto.Bet, err error)
	Update(ctx context.Context, bet dto.Bet) (updated bool, err error)
}
