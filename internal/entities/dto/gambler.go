package dto

import (
	"errors"
	"time"
)

var (
	empty_time = time.Time{}
)

type Gambler struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	Document     string    `json:"document"`
	DocumentType string    `json:"document_type"`
	BirthDate    time.Time `json:"birth_date"`
}

func (g *Gambler) Validate() error {
	if g.Name == "" {
		return errors.New("missing name")
	}

	if g.Document == "" {
		return errors.New("missing document")
	}

	if g.DocumentType == "" {
		return errors.New("missing document type")
	}

	if g.BirthDate == empty_time {
		return errors.New("missing birth date")
	}

	return nil
}
