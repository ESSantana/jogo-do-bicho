package entities

import "time"

type Bet struct {
	ID        int64      `db:"id" json:"id"`
	GamblerID int64      `db:"gambler_id" json:"gambler_id"`
	BetType   string     `db:"bet_type" json:"bet_type"`
	BetPrice  float64    `db:"bet_price" json:"bet_price"`
	BetChoice string     `db:"bet_choice" json:"bet_choice"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

func (b *Bet) IsValid() bool {
	return b.ID > 0 && b.GamblerID > 0
}