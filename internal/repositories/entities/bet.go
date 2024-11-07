package entities

import (
	"strconv"
	"strings"
	"time"
)

type BetType = string

const (
	Thousands    = "thousands"
	Hundreds     = "hundreds"
	Dozens       = "dozens"
	Group        = "group"
	DoubleDozens = "double_dozens"
	DoubleGroup  = "double_group"
)

type BetModifier = string

const (
	OnTop      = "on_top"
	Surrounded = "surrounded"
)

type Bet struct {
	ID             int64       `db:"id" json:"id"`
	GamblerID      int64       `db:"gambler_id" json:"gambler_id"`
	RaffleID       int64       `db:"raffle_id" json:"raffle_id"`
	BetType        BetType     `db:"bet_type" json:"bet_type"`
	BetModifier    BetModifier `db:"bet_modifier" json:"bet_modifier"`
	BetPrice       float64     `db:"bet_price" json:"bet_price"`
	BetCombination string      `db:"bet_combination" json:"bet_combination"`
	CreatedAt      *time.Time  `db:"created_at" json:"created_at"`
}

func (b *Bet) IsValid() bool {
	return b.ID > 0 && b.GamblerID > 0
}

func (b *Bet) GetCombinationIntValues() []int {
	var comb = make([]int, 0)
	for _, v := range strings.Split(b.BetCombination, ",") {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			return []int{}
		}
		comb = append(comb, intValue)
	}

	return comb
}
