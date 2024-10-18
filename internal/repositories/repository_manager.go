package repositories

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type RepositoryManager struct {
	queries *db.Queries
}

func NewRepositoryManager(ctx context.Context) *RepositoryManager {
	timeLoc, _ := time.LoadLocation("America/Sao_Paulo")
	cfg := mysql.Config{
		User:                 os.Getenv("db_user"),
		Passwd:               os.Getenv("db_pass"),
		Net:                  "tcp",
		Addr:                 os.Getenv("db_host"),
		DBName:               os.Getenv("db_name"),
		Loc:                  timeLoc,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	mysqlConn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	mysqlConn.SetConnMaxLifetime(time.Minute * 3)
	mysqlConn.SetMaxOpenConns(10)
	mysqlConn.SetMaxIdleConns(10)

	queries := db.New(mysqlConn)

	return &RepositoryManager{
		queries: queries,
	}
}

func (manager *RepositoryManager) NewBetRepository() contracts.BetRepository {
	return newBetRepository(manager.queries)
}

func (manager *RepositoryManager) NewGamblerRepository() contracts.GamblerRepository {
	return newGamblerRepository(manager.queries)
}
