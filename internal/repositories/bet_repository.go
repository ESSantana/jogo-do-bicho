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

func (r *BetRepository) test(ctx context.Context) {
	r.sqlcQueries.GetBet(ctx)
}
