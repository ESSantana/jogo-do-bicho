package entities

import "time"

type Raffle struct {
	ID            int64      `db:"id" json:"id"`
	RaffleEdition int32      `db:"raffle_edition" json:"raffle_edition"`
	Animal        string     `db:"animal" json:"animal"`
	RaffleNumber  string     `db:"raffle_number" json:"raffle_number"`
	RaffleOrder   int32      `db:"raffle_order" json:"raffle_order"`
	DeletedAt     *time.Time `db:"deleted_at" json:"deleted_at"`
}
