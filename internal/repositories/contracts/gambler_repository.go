package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
)

type GamblerRepository interface {
	Create(ctx context.Context, gambler db.CreateGamblerParams) (db.Gambler, error)
	GetAll(ctx context.Context) ([]db.GetGamblersRow, error)
	GetByID(ctx context.Context, id int32) (db.GetGamblerRow, error)
	// Update(ctx context.Context, gamblerUpdated db.UpdateGamblerParams) (db.Gambler, error)
	// Delete(ctx context.Context, deleteParams db.DeleteGamblerParams) (db.Gambler, error)
}
