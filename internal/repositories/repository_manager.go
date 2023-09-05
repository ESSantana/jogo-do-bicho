package repositories

import (
	"context"
	"database/sql"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
)

type RepositoryManager struct {
	sqlcQueries *db.Queries
}

func NewRepositoryManager(ctx context.Context) *RepositoryManager {
	dbconn, err := sql.Open("postgresql", "")
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	dbQueries := db.New(dbconn)

	return &RepositoryManager{
		sqlcQueries: dbQueries,
	}
}

func (manager *RepositoryManager) NewBetRepository() contracts.BetRepository {
	return newBetRepository(manager.sqlcQueries)
}
