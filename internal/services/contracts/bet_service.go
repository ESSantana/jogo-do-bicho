package contracts

import (
	"github.com/ESSantana/jogo-do-bicho/internal/common/entities/dto"
	vm "github.com/ESSantana/jogo-do-bicho/internal/common/entities/viewmodel"
)

type BetService interface {
	GetAll() []vm.Bet
	Get(id string) []vm.Bet
	Create(bet dto.Bet) vm.Bet
	Update(bet dto.Bet) vm.Bet
	Delete(id string) bool
}
