package entities

import "time"

type GamblersDocumentType = string

const (
	CPF GamblersDocumentType = "cpf"
	RG  GamblersDocumentType = "rg"
)

type Gambler struct {
	ID           int64                `db:"id" json:"id"`
	Name         string               `db:"name" json:"name"`
	Document     string               `db:"document" json:"document"`
	DocumentType GamblersDocumentType `db:"document_type" json:"document_type"`
	BirthDate    *time.Time           `db:"birth_date" json:"birth_date"`
	CreatedAt    *time.Time           `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time           `db:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time           `db:"deleted_at" json:"deleted_at"`
}

func (g *Gambler) IsValid() bool {
	return g.ID > 0
}
