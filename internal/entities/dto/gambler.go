package dto

import "time"

type Gambler struct {
	Name         string    `json:"name"`
	Document     string    `json:"document"`
	DocumentType string    `json:"document_type"`
	BirthDate    time.Time `json:"birth_date"`
}
