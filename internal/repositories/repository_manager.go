package repositories

import (
	"context"
	"os"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
	"github.com/jackc/pgx/v5"
)

type RepositoryManager struct {
	queries *db.Queries
}

func NewRepositoryManager(ctx context.Context) *RepositoryManager {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	queries := db.New(conn)

	return &RepositoryManager{
		queries: queries,
	}
}

func (manager *RepositoryManager) NewBetRepository() contracts.BetRepository {
	return newBetRepository(manager.queries)
}
