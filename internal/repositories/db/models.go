// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type GamblersDocumentType string

const (
	GamblersDocumentTypeCPF GamblersDocumentType = "CPF"
	GamblersDocumentTypeRG  GamblersDocumentType = "RG"
)

func (e *GamblersDocumentType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = GamblersDocumentType(s)
	case string:
		*e = GamblersDocumentType(s)
	default:
		return fmt.Errorf("unsupported scan type for GamblersDocumentType: %T", src)
	}
	return nil
}

type NullGamblersDocumentType struct {
	GamblersDocumentType GamblersDocumentType `json:"gamblers_document_type"`
	Valid                bool                 `json:"valid"` // Valid is true if GamblersDocumentType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGamblersDocumentType) Scan(value interface{}) error {
	if value == nil {
		ns.GamblersDocumentType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.GamblersDocumentType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGamblersDocumentType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.GamblersDocumentType), nil
}

type Bet struct {
	ID        int64        `db:"id" json:"id"`
	GamblerID int64        `db:"gambler_id" json:"gambler_id"`
	BetType   string       `db:"bet_type" json:"bet_type"`
	BetPrice  float64      `db:"bet_price" json:"bet_price"`
	BetChoice string       `db:"bet_choice" json:"bet_choice"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

type BetGroup struct {
	ID          int64  `db:"id" json:"id"`
	GroupName   string `db:"group_name" json:"group_name"`
	GroupNumber string `db:"group_number" json:"group_number"`
}

type Gambler struct {
	ID           int64                `db:"id" json:"id"`
	GamblerName  string               `db:"gambler_name" json:"gambler_name"`
	Document     string               `db:"document" json:"document"`
	DocumentType GamblersDocumentType `db:"document_type" json:"document_type"`
	BirthDate    time.Time            `db:"birth_date" json:"birth_date"`
	UpdatedAt    sql.NullTime         `db:"updated_at" json:"updated_at"`
	DeletedAt    sql.NullTime         `db:"deleted_at" json:"deleted_at"`
}

type Raffle struct {
	ID            int64        `db:"id" json:"id"`
	RaffleEdition int32        `db:"raffle_edition" json:"raffle_edition"`
	Animal        string       `db:"animal" json:"animal"`
	RaffleNumber  string       `db:"raffle_number" json:"raffle_number"`
	RaffleOrder   int32        `db:"raffle_order" json:"raffle_order"`
	DeletedAt     sql.NullTime `db:"deleted_at" json:"deleted_at"`
}
