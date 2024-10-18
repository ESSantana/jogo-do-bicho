package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ESSantana/jogo-do-bicho/internal/repositories/contracts"
	"github.com/ESSantana/jogo-do-bicho/internal/repositories/entities"
)

type BetRepository struct {
	conn *sql.DB
}

func newBetRepository(conn *sql.DB) contracts.BetRepository {
	return &BetRepository{
		conn: conn,
	}
}

func (repo *BetRepository) Create(ctx context.Context, bet entities.Bet) (persistedID int64, err error) {
	result, err := repo.conn.ExecContext(ctx, `
		INSERT INTO
			bets (gambler_id, bet_type, bet_price, bet_choice)
		VALUES
			(?, ?, ?, ?);
	`,
		bet.GamblerID, bet.BetType, bet.BetPrice, bet.BetChoice,
	)

	if err != nil {
		return persistedID, err
	}

	if affected, err := result.RowsAffected(); affected == 1 && err == nil {
		return result.LastInsertId()
	}

	return persistedID, errors.New("nenhuma linha foi afetada")
}

func (repo *BetRepository) GetAll(ctx context.Context) (bets []entities.Bet, err error) {
	rows, err := repo.conn.QueryContext(ctx, `
		SELECT
			*
		FROM
			bets
		WHERE
			bets.deleted_at IS NULL
		`)

	if err != nil {
		return bets, err
	}

	defer rows.Close()
	bets = []entities.Bet{}

	for rows.Next() {
		var bet entities.Bet
		err := rows.Scan(
			&bet.ID,
			&bet.GamblerID,
			&bet.BetType,
			&bet.BetPrice,
			&bet.BetChoice,
			&bet.DeletedAt,
		)
		if err != nil {
			return bets, err
		}

		bets = append(bets, bet)
	}

	if err = rows.Err(); err != nil {
		return bets, err
	}

	return bets, nil
}

func (repo *BetRepository) GetByID(ctx context.Context, id int64) (bet entities.Bet, err error) {
	row := repo.conn.QueryRowContext(ctx, `
		SELECT
			*
		FROM
			bets
		WHERE
			bets.id = ?
			AND bets.deleted_at IS NULL
		`,
		id,
	)

	err = row.Scan(
		&bet.ID,
		&bet.GamblerID,
		&bet.BetType,
		&bet.BetPrice,
		&bet.BetChoice,
		&bet.DeletedAt,
	)

	return bet, err
}

func (repo *BetRepository) Update(ctx context.Context, bet entities.Bet) (rowsAffected int64, err error) {
	result, err := repo.conn.ExecContext(ctx, `
		UPDATE
			bets
		SET
			bet_type = ?,
			bet_price = ?,
			bet_choice = ?
		WHERE
			id = ?
			AND bets.deleted_at IS NULL
		`,
		bet.BetType, bet.BetPrice, bet.BetChoice, bet.ID,
	)

	if err != nil {
		return rowsAffected, err
	}

	return result.RowsAffected()
}

func (repo *BetRepository) Delete(ctx context.Context, bet entities.Bet) (rowsAffected int64, err error) {
	result, err := repo.conn.ExecContext(ctx, `
		UPDATE
			bets
		SET
			deleted_at = ?
		WHERE
			id = ?
		`,
		bet.DeletedAt, bet.ID,
	)

	if err != nil {
		return rowsAffected, err
	}

	return result.RowsAffected()
}
