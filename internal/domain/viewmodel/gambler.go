package viewmodel

import "time"

type Gambler struct {
	ID           int64     `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Document     string    `json:"document,omitempty"`
	DocumentType string    `json:"document_type,omitempty"`
	BirthDate    time.Time `json:"birth_date,omitempty"`
	Bets         []Bet     `json:"bets,omitempty"`
}
