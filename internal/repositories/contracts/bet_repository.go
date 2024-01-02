package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
)

type BetRepository interface {
	Create(ctx context.Context, bet db.CreateBetParams) (db.Bet, error)
	GetAll(ctx context.Context) ([]db.GetBetsRow, error)
	GetByID(ctx context.Context, id int32) (db.GetBetRow, error)
	Update(ctx context.Context, betUpdated db.UpdateBetParams) (db.Bet, error)
	Delete(ctx context.Context, deleteParams db.DeleteBetParams) (db.Bet, error)
}
