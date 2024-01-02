// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type DocType string

const (
	DocTypeCPF DocType = "CPF"
	DocTypeRG  DocType = "RG"
)

func (e *DocType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = DocType(s)
	case string:
		*e = DocType(s)
	default:
		return fmt.Errorf("unsupported scan type for DocType: %T", src)
	}
	return nil
}

type NullDocType struct {
	DocType DocType `json:"doc_type"`
	Valid   bool    `json:"valid"` // Valid is true if DocType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullDocType) Scan(value interface{}) error {
	if value == nil {
		ns.DocType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.DocType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullDocType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.DocType), nil
}

type Bet struct {
	ID        int32            `db:"id" json:"id"`
	GamblerID int32            `db:"gambler_id" json:"gambler_id"`
	BetType   string           `db:"bet_type" json:"bet_type"`
	BetPrice  float64          `db:"bet_price" json:"bet_price"`
	BetChoice string           `db:"bet_choice" json:"bet_choice"`
	DeletedAt pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type Gambler struct {
	ID           int32   `db:"id" json:"id"`
	GamblerName  string  `db:"gambler_name" json:"gambler_name"`
	Document     string  `db:"document" json:"document"`
	DocumentType DocType `db:"document_type" json:"document_type"`
	BirthDate    string  `db:"birth_date" json:"birth_date"`
}

type Raffle struct {
	ID            int32  `db:"id" json:"id"`
	RaffleEdition int32  `db:"raffle_edition" json:"raffle_edition"`
	Animal        string `db:"animal" json:"animal"`
	RaffleNumber  string `db:"raffle_number" json:"raffle_number"`
	RaffleOrder   int32  `db:"raffle_order" json:"raffle_order"`
}
