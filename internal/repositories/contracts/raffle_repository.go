package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/entities"
)

type RaffleRepository interface {
	Create(ctx context.Context, raffle entities.Raffle) (persistedID int64, err error)
	GetAll(ctx context.Context) (raffles []entities.Raffle, err error)
	GetByID(ctx context.Context, id int64) (raffle entities.Raffle, err error)
	GetCurrentEdition(ctx context.Context) (raffle entities.Raffle, err error)
	GetAllDrawsByRaffleID(ctx context.Context, raffleID int64) (raffles []entities.RaffleDraw, err error)
	Update(ctx context.Context, raffle entities.Raffle) (rowsAffected int64, err error)
	Delete(ctx context.Context, raffle entities.Raffle) (rowsAffected int64, err error)
}
