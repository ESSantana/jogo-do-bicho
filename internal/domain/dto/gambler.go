package dto

import (
	"errors"
	"time"

	custom_errors "github.com/ESSantana/jogo-do-bicho/internal/domain/errors"
)

var (
	empty_time = time.Time{}
)

type GamblersDocumentType = string

const (
	CPF GamblersDocumentType = "cpf"
	RG  GamblersDocumentType = "rg"
)

type Gambler struct {
	ID           int64                `json:"id,omitempty"`
	Name         string               `json:"name"`
	Document     string               `json:"document"`
	DocumentType GamblersDocumentType `json:"document_type"`
	BirthDate    time.Time            `json:"birth_date"`
	Bets         []Bet                `json:"bets,omitempty"`
	CreatedAt    *time.Time           `json:"created_at,omitempty"`
	UpdatedAt    *time.Time           `json:"updated_at,omitempty"`
}

func (g *Gambler) Validate() error {
	var errs []error

	if g.Name == "" {
		errs = append(errs, custom_errors.NewValidationError("campo nome é obrigatório"))
	}

	if g.Document == "" {
		errs = append(errs, custom_errors.NewValidationError("campo documento é obrigatório"))
	}

	if g.DocumentType != CPF && g.DocumentType != RG {
		errs = append(errs, custom_errors.NewValidationError("campo tipo de documento é inválido"))
	}

	if g.BirthDate == empty_time {
		errs = append(errs, custom_errors.NewValidationError("campo data de nascimento é obrigatório"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
