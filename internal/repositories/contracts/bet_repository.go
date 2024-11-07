package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/entities"
)

type BetRepository interface {
	Create(ctx context.Context, bet entities.Bet) (persistedID int64, err error)
	GetAll(ctx context.Context) (bets []entities.Bet, err error)
	GetAllByGamblerID(ctx context.Context, gamblerID int64) (bets []entities.Bet, err error)
	GetByID(ctx context.Context, id int64) (bet entities.Bet, err error)
	Update(ctx context.Context, bet entities.Bet) (rowsAffected int64, err error)
}
