package dto

import "time"

type PrizeDraw struct {
	Order    int        `json:"order"`
	Number   int        `json:"number"`
	OccursAt *time.Time `json:"occurs_at"`
}

type Raffle struct {
	ID         int         `json:"id"`
	Edition    string      `json:"edition"`
	PrizeDraws []PrizeDraw `json:"prize_draws,omitempty"`
	CreatedAt  *time.Time  `json:"created_at,omitempty"`
	UpdatedAt  *time.Time  `json:"updated_at,omitempty"`
	DeletedAt  *time.Time  `json:"deleted_at,omitempty"`
}
