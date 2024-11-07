package dto

import (
	"errors"
	"time"
)

type BetTypeSlug string

const (
	Thousands    BetTypeSlug = "thousands"
	Hundreds     BetTypeSlug = "hundreds"
	Dozens       BetTypeSlug = "dozens"
	Group        BetTypeSlug = "group"
	DoubleDozens BetTypeSlug = "double_dozens"
	DoubleGroup  BetTypeSlug = "double_group"
)

var betSlugs = []BetTypeSlug{"thousands", "hundreds", "dozens", "group", "double_dozens", "double_group"}

func (bts BetTypeSlug) GetPrizeMultiplier() float64 {
	switch bts {
	case Thousands:
		return 4000
	case Hundreds:
		return 600
	case Dozens:
		return 60
	case Group:
		return 18
	case DoubleDozens:
		return 300
	case DoubleGroup:
		return 18.75
	default:
		return 0.0
	}
}

type BetModifier = string

const (
	OnTop      BetModifier = "on_top"
	Surrounded BetModifier = "surrounded"
)

type BetType struct {
	Slug         BetTypeSlug `json:"slug"`
	FriendlyName string      `json:"name"`
	Description  string      `json:"description"`
}

func (bt *BetType) Validate() error {
	validBetType := false
	for _, val := range betSlugs {
		if bt.Slug == val {
			validBetType = true
		}
	}

	if !validBetType {
		return errors.New("tipo de aposta deve ter identificador válido")
	}

	return nil
}

type BetGroup struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	GroupNumbers []int  `json:"group_numbers"`
}

func (bg *BetGroup) IsNumberPartOfGroup(input int) bool {
	for _, n := range bg.GroupNumbers {
		if n == input {
			return true
		}
	}
	return false
}

type Bet struct {
	ID             int64       `json:"id,omitempty"`
	RaffleID       int64       `json:"raffle_id"`
	GamblerID      int64       `json:"gambler_id"`
	BetPrice       float64     `json:"bet_price"`
	BetType        BetType     `json:"bet_type"`
	BetModifier    BetModifier `json:"bet_modifier"`
	BetCombination []int       `json:"bet_combination"`
	CreatedAt      time.Time   `json:"created_at"`
}

//TODO: implementar validação de raffleID
func (b *Bet) Validate() error {
	var errorsValidation = []error{}

	if b.GamblerID <= 0 {
		errorsValidation = append(errorsValidation, errors.New("aposta deve ser atribuída a um jogador válido"))
	}

	if b.BetPrice <= 0 {
		errorsValidation = append(errorsValidation, errors.New("valor da aposta deve ser maior que 0"))
	}

	if btError := b.BetType.Validate(); btError != nil {
		errorsValidation = append(errorsValidation, btError)
	}

	if (b.BetModifier != OnTop && b.BetModifier != Surrounded) && (b.BetType.Slug != DoubleDozens && b.BetType.Slug != DoubleGroup) {
		errorsValidation = append(errorsValidation, errors.New("modificador de aposta inválido"))
	}

	if len(b.BetCombination) > 1 && b.BetType.Slug != DoubleDozens && b.BetType.Slug != DoubleGroup {
		errorsValidation = append(errorsValidation, errors.New("aposta com mais de um número não é válida"))
	}

	if len(errorsValidation) > 0 {
		return errors.Join(errorsValidation...)
	}

	for _, combination := range b.BetCombination {
		if b.BetType.Slug == DoubleGroup && (combination < 1 || combination > 25) {
			errorsValidation = append(errorsValidation, errors.New("grupo selecionado na aposta inválido"))
			continue
		}

		if b.BetType.Slug == DoubleDozens && (combination < 0 || combination > 99) {
			errorsValidation = append(errorsValidation, errors.New("dezena selecionado na aposta inválido"))
			continue
		}
	}

	if len(errorsValidation) > 0 {
		return errors.Join(errorsValidation...)
	}

	return nil
}
