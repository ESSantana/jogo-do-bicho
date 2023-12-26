package services

import (
	"context"

	"github.com/ESSantana/jogo-do-bicho/internal/entities/dto"
	vm "github.com/ESSantana/jogo-do-bicho/internal/entities/viewmodel"
	repo_contracts "github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/db"
	"github.com/ESSantana/jogo-do-bicho/internal/services/contracts"
)

type BetService struct {
	repoManager repo_contracts.RepositoryManager
}

func newBetService(repoManager repo_contracts.RepositoryManager) contracts.BetService {
	return &BetService{
		repoManager: repoManager,
	}
}

func (svc *BetService) Create(ctx context.Context, bet dto.Bet) (createdBet vm.Bet, err error) {
	betParams := db.CreateBetParams{
		GamblerID: int32(bet.GamblerID),
		BetType:   bet.BetType,
		BetPrice:  bet.BetPrice,
		BetChoice: bet.BetChoice,
	}

	betRepo := svc.repoManager.NewBetRepository()
	persistedBet, err := betRepo.CreateBet(ctx, betParams)
	if err != nil {
		return createdBet, err
	}

	returnBet := vm.Bet{
		ID: int(createdBet.ID),
		Gambler: vm.Gambler{
			ID: persistedBet.GamblerID,
		},
		BetType:   createdBet.BetType,
		BetPrice:  createdBet.BetPrice,
		BetChoice: createdBet.BetChoice,
	}

	return returnBet, err
}

func (svc *BetService) GetAllBets(ctx context.Context) (allBets []vm.Bet, err error) {
	betRepo := svc.repoManager.NewBetRepository()
	items, err := betRepo.GetAllBets(ctx)
	if err != nil {
		return nil, err
	}

	allBets = make([]vm.Bet, 0)
	for _, item := range items {
		allBets = append(allBets, vm.Bet{
			ID:        int(item.Bet.ID),
			BetType:   item.Bet.BetType,
			BetPrice:  item.Bet.BetPrice,
			BetChoice: item.Bet.BetChoice,
			Gambler: vm.Gambler{
				ID:   item.Gambler.ID,
				Name: item.Gambler.GamblerName,
			},
		})
	}
	return allBets, nil
}

func (svc *BetService) GetByID(ctx context.Context, id int32) (bet vm.Bet, err error) {
	betRepo := svc.repoManager.NewBetRepository()
	betPersisted, err := betRepo.GetByID(ctx, id)
	if err != nil {
		return bet, err
	}

	bet = vm.Bet{
		ID:        int(betPersisted.Bet.ID),
		BetType:   betPersisted.Bet.BetType,
		BetPrice:  betPersisted.Bet.BetPrice,
		BetChoice: betPersisted.Bet.BetChoice,
		Gambler: vm.Gambler{
			ID:   betPersisted.Gambler.ID,
			Name: betPersisted.Gambler.GamblerName,
		},
	}

	return bet, nil
}

// func (svc *BetService) Update(ctx context.Context) (allBets []vm.Bet, err error) {
// 	betRepo := svc.repoManager.NewBetRepository()
// 	items, err := betRepo.GetAllBets(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	allBets = make([]vm.Bet, 0)
// 	for _, item := range items {
// 		allBets = append(allBets, vm.Bet{
// 			ID:        int(item.Bet.ID),
// 			BetType:   item.Bet.BetType,
// 			BetPrice:  item.Bet.BetPrice,
// 			BetChoice: item.Bet.BetChoice,
// 			Gambler: vm.Gambler{
// 				ID:   item.Gambler.ID,
// 				Name: item.Gambler.GamblerName,
// 			},
// 		})
// 	}
// 	return allBets, nil
// }

// func (svc *BetService) Delete(ctx context.Context) (allBets []vm.Bet, err error) {
// 	betRepo := svc.repoManager.NewBetRepository()
// 	items, err := betRepo.GetAllBets(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	allBets = make([]vm.Bet, 0)
// 	for _, item := range items {
// 		allBets = append(allBets, vm.Bet{
// 			ID:        int(item.Bet.ID),
// 			BetType:   item.Bet.BetType,
// 			BetPrice:  item.Bet.BetPrice,
// 			BetChoice: item.Bet.BetChoice,
// 			Gambler: vm.Gambler{
// 				ID:   item.Gambler.ID,
// 				Name: item.Gambler.GamblerName,
// 			},
// 		})
// 	}
// 	return allBets, nil
// }
