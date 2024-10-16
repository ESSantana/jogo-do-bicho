package contracts

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/domain/dto"
	vm "github.com/ESSantana/jogo-do-bicho/internal/domain/viewmodel"
)

type GamblerService interface {
	GetAll(ctx context.Context) (gamblers []vm.Gambler, err error)
	GetByID(ctx context.Context, id int64) (gambler vm.Gambler, err error)
	Create(ctx context.Context, gambler dto.Gambler) (gamblerVm vm.Gambler, err error)
	Update(ctx context.Context, gambler dto.Gambler) (updated bool, err error)
	Delete(ctx context.Context, id int64) (deleted bool, err error)
}
