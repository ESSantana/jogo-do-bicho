package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
)

type BetRepository interface {
	CreateBet(ctx context.Context, bet db.CreateBetParams) (db.Bet, error)
	GetAllBets(ctx context.Context) ([]db.GetBetsRow, error)
}
