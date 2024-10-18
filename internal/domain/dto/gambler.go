package dto

import (
	"errors"
	"time"

	custom_errors "github.com/ESSantana/jogo-do-bicho/internal/domain/errors"
)

var (
	empty_time = time.Time{}
)

type Gambler struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Document     string    `json:"document"`
	DocumentType string    `json:"document_type"`
	BirthDate    time.Time `json:"birth_date"`
}

func (g *Gambler) Validate() error {
	var errs []error

	if g.Name == "" {
		errs = append(errs, custom_errors.NewValidationError("campo nome é obrigatório"))
	}

	if g.Document == "" {
		errs = append(errs, custom_errors.NewValidationError("campo documento é obrigatório"))
	}

	if g.DocumentType == "" {
		errs = append(errs, custom_errors.NewValidationError("campo tipo de documento é obrigatório"))
	}

	if g.BirthDate == empty_time {
		errs = append(errs, custom_errors.NewValidationError("campo data de nascimento é obrigatório"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
