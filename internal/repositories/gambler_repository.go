package repositories

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
)

type GamblerRepository struct {
	sqlcQueries *db.Queries
}

func newGamblerRepository(sqlcQueries *db.Queries) contracts.GamblerRepository {
	return &GamblerRepository{
		sqlcQueries: sqlcQueries,
	}
}

func (r *GamblerRepository) Create(ctx context.Context, gambler db.CreateGamblerParams) (error) {
	err := r.sqlcQueries.CreateGambler(ctx, gambler)
	return err
}

func (r *GamblerRepository) GetAll(ctx context.Context) ([]db.GetGamblersRow, error) {
	createdGambler, err := r.sqlcQueries.GetGamblers(ctx)
	return createdGambler, err
}

func (r *GamblerRepository) GetByID(ctx context.Context, id int64) ([]db.GetGamblerRow, error) {
	gambler, err := r.sqlcQueries.GetGambler(ctx, id)
	return gambler, err
}

func (r *GamblerRepository) Update(ctx context.Context, gamblerUpdated db.UpdateGamblerParams) (error) {
	 err := r.sqlcQueries.UpdateGambler(ctx, gamblerUpdated)
	return  err
}

func (r *GamblerRepository) Delete(ctx context.Context, deleteParams db.DeleteGamblerParams) (error) {
	 err := r.sqlcQueries.DeleteGambler(ctx, deleteParams)
	return  err
}
