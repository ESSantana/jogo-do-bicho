package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/entities"
)

type GamblerRepository struct {
	conn *sql.DB
}

func newGamblerRepository(conn *sql.DB) contracts.GamblerRepository {
	return &GamblerRepository{
		conn: conn,
	}
}

func (repo *GamblerRepository) Create(ctx context.Context, gambler entities.Gambler) (persistedID int64, err error) {
	result, err := repo.conn.ExecContext(ctx, `
		INSERT INTO
			gambler (name, document, document_type, birth_date, updated_at)
		VALUES
			(?, ?, ?, ?, ?)
	`,
		gambler.Name, gambler.Document, gambler.DocumentType, gambler.BirthDate, gambler.UpdatedAt,
	)

	if err != nil {
		return persistedID, err
	}

	if affected, err := result.RowsAffected(); affected == 1 && err == nil {
		return result.LastInsertId()
	}

	return persistedID, errors.New("nenhuma linha foi afetada")
}

func (repo *GamblerRepository) GetAll(ctx context.Context) (gamblers []entities.Gambler, err error) {
	rows, err := repo.conn.QueryContext(ctx, `
		SELECT
			*
		FROM
			gambler
		WHERE
			deleted_at IS NULL
		`)

	if err != nil {
		return gamblers, err
	}

	defer rows.Close()
	gamblers = []entities.Gambler{}

	for rows.Next() {
		var gambler entities.Gambler
		err := rows.Scan(
			&gambler.ID,
			&gambler.Name,
			&gambler.Document,
			&gambler.DocumentType,
			&gambler.BirthDate,
			&gambler.CreatedAt,
			&gambler.UpdatedAt,
			&gambler.DeletedAt,
		)
		if err != nil {
			return gamblers, err
		}

		gamblers = append(gamblers, gambler)
	}

	if err = rows.Err(); err != nil {
		return gamblers, err
	}

	return gamblers, nil
}

func (repo *GamblerRepository) GetByID(ctx context.Context, id int64) (gambler entities.Gambler, err error) {
	row := repo.conn.QueryRowContext(ctx, `
		SELECT
			*
		FROM
			gambler
		WHERE
			gambler.id = ?
			AND gambler.deleted_at IS NULL
		`,
		id,
	)

	err = row.Scan(
		&gambler.ID,
		&gambler.Name,
		&gambler.Document,
		&gambler.DocumentType,
		&gambler.BirthDate,
		&gambler.CreatedAt,
		&gambler.UpdatedAt,
		&gambler.DeletedAt,
	)

	return gambler, err
}

func (repo *GamblerRepository) Update(ctx context.Context, gambler entities.Gambler) (rowsAffected int64, err error) {
	result, err := repo.conn.ExecContext(ctx, `
		UPDATE
			gambler
		SET
			name = ?,
			document = ?,
			document_type = ?,
			birth_date = ?
			updated_at = ?,
		WHERE
			id = ?
			AND gambler.deleted_at IS NULL
		`,
		gambler.Name, gambler.Document, gambler.DocumentType, gambler.BirthDate, gambler.UpdatedAt, gambler.ID,
	)

	if err != nil {
		return rowsAffected, err
	}

	return result.RowsAffected()
}

func (repo *GamblerRepository) Delete(ctx context.Context, gambler entities.Gambler) (rowsAffected int64, err error) {
	result, err := repo.conn.ExecContext(ctx, `
		UPDATE
			gambler
		SET
			deleted_at = ?
		WHERE
			id = ?
		`,
		gambler.DeletedAt, gambler.ID,
	)

	if err != nil {
		return rowsAffected, err
	}

	return result.RowsAffected()
}
