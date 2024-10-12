package dto

import "errors"

type BetTypeSlug string

const (
	Thousands    BetTypeSlug = "thousands"
	Hundreds     BetTypeSlug = "hundreds"
	Dozens       BetTypeSlug = "dozens"
	Group        BetTypeSlug = "group"
	DoubleDozens BetTypeSlug = "double_dozens"
	DoubleGroup  BetTypeSlug = "double_group"
)

var betSlugArr = []BetTypeSlug{"thousands", "hundreds", "dozens", "group", "double_dozens", "double_group"}

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

type BetModifier string

const (
	OnTop      BetModifier = "on_top"
	Surrounded BetModifier = "surrounded"
)

type BetType struct {
	ID          int         `json:"id"`
	Slug        BetTypeSlug `json:"slug"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
}

func (bt *BetType) Validate() error {
	var errorsValidation = []error{}

	if bt.ID <= 0 {
		errorsValidation = append(errorsValidation, errors.New("tipo de aposta deve ser um valor válido"))
	}

	validBetType := false
	for _, val := range betSlugArr {

		if bt.Slug == val {
			validBetType = true
		}
	}

	if !validBetType {
		errorsValidation = append(errorsValidation, errors.New("tipo de aposta deve ter identificador válido"))
	}

	if len(errorsValidation) > 0 {
		return errors.Join(errorsValidation...)
	}

	return nil
}

type BetGroup struct {
	ID           int    `json:"id"`
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
	ID             int         `json:"id,omitempty"`
	GamblerID      int         `json:"gambler_id"`
	BetType        BetType     `json:"bet_type"`
	BetModifier    BetModifier `json:"bet_modifier"`
	BetPrice       float64     `json:"bet_price"`
	BetCombination []int       `json:"bet_combination"`
}

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

	if b.BetModifier != OnTop || b.BetModifier != Surrounded {
		errorsValidation = append(errorsValidation, errors.New("modificador de aposta inválido"))
	}

	for _, combination := range b.BetCombination {
		if b.BetType.Slug == Group && (combination < 1 || combination > 25) {
			errorsValidation = append(errorsValidation, errors.New("grupo selecionado na aposta inválido"))
			break
		}
		
		if b.BetType.Slug == Group && (combination < 1 || combination > 25) {
			errorsValidation = append(errorsValidation, errors.New("grupo selecionado na aposta inválido"))
			break
		}

		if b.BetType.Slug == Group && (combination < 1 || combination > 25) {
			errorsValidation = append(errorsValidation, errors.New("grupo selecionado na aposta inválido"))
			break
		}

		if b.BetType.Slug == Group && (combination < 1 || combination > 25) {
			errorsValidation = append(errorsValidation, errors.New("grupo selecionado na aposta inválido"))
			break
		}
	}

	if len(errorsValidation) > 0 {
		return errors.Join(errorsValidation...)
	}

	return nil
}
