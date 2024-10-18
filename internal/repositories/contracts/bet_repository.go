package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
)

type BetRepository interface {
	Create(ctx context.Context, bet db.CreateBetParams) (error)
	GetAll(ctx context.Context) ([]db.GetBetsRow, error)
	GetByID(ctx context.Context, id int64) (db.GetBetRow, error)
	Update(ctx context.Context, betUpdated db.UpdateBetParams) (error)
	Delete(ctx context.Context, deleteParams db.DeleteBetParams) (error)
}
