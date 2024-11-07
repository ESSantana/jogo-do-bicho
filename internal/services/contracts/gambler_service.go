package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/domain/dto"
)

type GamblerService interface {
	GetAll(ctx context.Context) (gamblers []dto.Gambler, err error)
	GetByID(ctx context.Context, id int64) (gambler dto.Gambler, err error)
	Create(ctx context.Context, gambler dto.Gambler) (gamblerOut dto.Gambler, err error)
	Update(ctx context.Context, gambler dto.Gambler) (updated bool, err error)
	Delete(ctx context.Context, id int64) (deleted bool, err error)
}
