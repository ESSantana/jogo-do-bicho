package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/entities"
)

type GamblerRepository interface {
	Create(ctx context.Context, gambler entities.Gambler) (persistedID int64, err error)
	GetAll(ctx context.Context) (gamblers []entities.Gambler, err error)
	GetByID(ctx context.Context, id int64) (gambler entities.Gambler, err error)
	Update(ctx context.Context, gambler entities.Gambler) (rowsAffected int64, err error)
	Delete(ctx context.Context, gambler entities.Gambler) (rowsAffected int64, err error)
}
