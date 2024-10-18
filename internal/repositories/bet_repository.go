package repositories

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
)

type BetRepository struct {
	sqlcQueries *db.Queries
}

func newBetRepository(sqlcQueries *db.Queries) contracts.BetRepository {
	return &BetRepository{
		sqlcQueries: sqlcQueries,
	}
}

func (r *BetRepository) Create(ctx context.Context, bet db.CreateBetParams) ( error) {
	 err := r.sqlcQueries.CreateBet(ctx, bet)
	return err
}

func (r *BetRepository) GetAll(ctx context.Context) ([]db.GetBetsRow, error) {
	createdBet, err := r.sqlcQueries.GetBets(ctx)
	return createdBet, err
}

func (r *BetRepository) GetByID(ctx context.Context, id int64) (db.GetBetRow, error) {
	bet, err := r.sqlcQueries.GetBet(ctx, id)
	return bet, err
}

func (r *BetRepository) Update(ctx context.Context, betUpdated db.UpdateBetParams) (error) {
	err := r.sqlcQueries.UpdateBet(ctx, betUpdated)
	return err
}

func (r *BetRepository) Delete(ctx context.Context, deleteParams db.DeleteBetParams) (error) {
	err := r.sqlcQueries.DeleteBet(ctx, deleteParams)
	return err
}
