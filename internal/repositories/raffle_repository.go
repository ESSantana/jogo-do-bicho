package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/entities"
)

type raffleRepository struct {
	conn *sql.DB
}

func newRaffleRepository(conn *sql.DB) contracts.RaffleRepository {
	return &raffleRepository{
		conn: conn,
	}
}

func (repo *raffleRepository) Create(ctx context.Context, raffle entities.Raffle) (persistedID int64, err error) {
	result, err := repo.conn.ExecContext(ctx, `
		INSERT INTO
			raffle (
				edition,
				updated_at,
			)
		VALUES
			(?, ?);
	`,
		raffle.Edition, raffle.UpdatedAt,
	)

	if err != nil {
		return persistedID, err
	}

	if affected, err := result.RowsAffected(); affected == 1 && err == nil {
		return result.LastInsertId()
	}

	return persistedID, errors.New("nenhuma linha foi afetada")
}

func (repo *raffleRepository) GetAll(ctx context.Context) (raffles []entities.Raffle, err error) {
	rows, err := repo.conn.QueryContext(ctx, `
		SELECT
			*
		FROM
			raffle
			AND deleted_at IS NULL
		`)

	if err != nil {
		return raffles, err
	}

	defer rows.Close()
	raffles = []entities.Raffle{}

	for rows.Next() {
		var raffle entities.Raffle
		err := rows.Scan(
			&raffle.ID,
			&raffle.Edition,
			&raffle.CreatedAt,
			&raffle.UpdatedAt,
			&raffle.DeletedAt,
		)
		if err != nil {
			return raffles, err
		}

		raffles = append(raffles, raffle)
	}

	if err = rows.Err(); err != nil {
		return raffles, err
	}

	return raffles, nil
}

func (repo *raffleRepository) GetAllDrawsByRaffleID(ctx context.Context, raffleID int64) (raffleDraws []entities.RaffleDraw, err error) {
	rows, err := repo.conn.QueryContext(ctx, `
		SELECT
			*
		FROM
			raffle_draw
		WHERE
			raffle.id = ?
		`,
		raffleID,
	)

	if err != nil {
		return raffleDraws, err
	}

	defer rows.Close()
	raffleDraws = []entities.RaffleDraw{}

	for rows.Next() {
		var rd entities.RaffleDraw
		err := rows.Scan(
			&rd.ID,
			&rd.RaffleID,
			&rd.Order,
			&rd.Number,
			&rd.CreatedAt,
		)
		if err != nil {
			return raffleDraws, err
		}

		raffleDraws = append(raffleDraws, rd)
	}

	if err = rows.Err(); err != nil {
		return raffleDraws, err
	}

	return raffleDraws, nil
}

func (repo *raffleRepository) GetByID(ctx context.Context, id int64) (raffle entities.Raffle, err error) {
	row := repo.conn.QueryRowContext(ctx, `
		SELECT
			*
		FROM
			raffle
		WHERE
			raffle.id = ?
			AND deleted_at IS NULL
		`,
		id,
	)

	err = row.Scan(
		&raffle.ID,
		&raffle.Edition,
		&raffle.CreatedAt,
		&raffle.UpdatedAt,
		&raffle.DeletedAt,
	)

	return raffle, err
}

func (repo *raffleRepository) GetCurrentEdition(ctx context.Context) (raffle entities.Raffle, err error) {
	row := repo.conn.QueryRowContext(ctx, `SELECT * FROM raffle WHERE deleted_at IS NULL ORDER BY created_at LIMIT 1`)

	err = row.Scan(
		&raffle.ID,
		&raffle.Edition,
		&raffle.CreatedAt,
		&raffle.UpdatedAt,
		&raffle.DeletedAt,
	)

	return raffle, err
}

func (repo *raffleRepository) Update(ctx context.Context, raffle entities.Raffle) (rowsAffected int64, err error) {
	result, err := repo.conn.ExecContext(ctx, `
		UPDATE
			raffle
		SET
			edition = ?,
		WHERE
			id = ?
			AND deleted_at IS NULL
		`,
		raffle.Edition, raffle.ID,
	)

	if err != nil {
		return rowsAffected, err
	}

	return result.RowsAffected()
}

func (repo *raffleRepository) Delete(ctx context.Context, raffle entities.Raffle) (rowsAffected int64, err error) {
	result, err := repo.conn.ExecContext(ctx, `
		UPDATE
			raffle
		SET
			deleted_at = ?
		WHERE
			id = ?
			AND deleted_at IS NULL
		`,
		raffle.DeletedAt, raffle.ID,
	)

	if err != nil {
		return rowsAffected, err
	}

	return result.RowsAffected()
}
