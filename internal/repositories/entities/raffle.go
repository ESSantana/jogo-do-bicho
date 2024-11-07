package entities

import "time"

type Raffle struct {
	ID        int64      `db:"id" json:"id"`
	Edition   int        `db:"edition" json:"edition"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type RaffleDraw struct {
	ID        int64      `db:"id" json:"id"`
	RaffleID  int64      `db:"raffle_id" json:"raffle_id"`
	Order     int        `db:"order" json:"order"`
	Number    string     `db:"number" json:"number"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}
