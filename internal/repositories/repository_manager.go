package repositories

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/go-sql-driver/mysql"
)

type RepositoryManager struct {
	conn *sql.DB
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

	return &RepositoryManager{
		conn: mysqlConn,
	}
}

func (manager *RepositoryManager) NewBetRepository() contracts.BetRepository {
	return newBetRepository(manager.conn)
}

func (manager *RepositoryManager) NewGamblerRepository() contracts.GamblerRepository {
	return newGamblerRepository(manager.conn)
}

func (manager *RepositoryManager) NewRaffleRepository() contracts.RaffleRepository {
	return newRaffleRepository(manager.conn)
}
