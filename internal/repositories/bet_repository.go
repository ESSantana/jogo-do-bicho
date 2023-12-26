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

func (r *BetRepository) CreateBet(ctx context.Context, bet db.CreateBetParams) (db.Bet, error) {
	createdBet, err := r.sqlcQueries.CreateBet(ctx, bet)
	return createdBet, err
}

func (r *BetRepository) GetAllBets(ctx context.Context) ([]db.GetBetsRow, error) {
	createdBet, err := r.sqlcQueries.GetBets(ctx)
	return createdBet, err
}

func (r *BetRepository) GetByID(ctx context.Context, id int32) (db.GetBetRow, error) {
	bet, err := r.sqlcQueries.GetBet(ctx, id)
	return bet, err
}
