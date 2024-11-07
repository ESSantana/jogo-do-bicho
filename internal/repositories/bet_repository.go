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
			bet (
				gambler_id,
				raffle_id,
				bet_type,
				bet_modifier,
				bet_price,
				bet_combination,
				created_at
			)
		VALUES
			(?, ?, ?, ?, ?, ?, ?);
	`,
		bet.GamblerID, bet.RaffleID, bet.BetType, bet.BetModifier, bet.BetPrice, bet.BetCombination, bet.CreatedAt,
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
			bet
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
			&bet.RaffleID,
			&bet.BetType,
			&bet.BetModifier,
			&bet.BetPrice,
			&bet.BetCombination,
			&bet.CreatedAt,
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

func (repo *BetRepository) GetAllByGamblerID(ctx context.Context, gamblerID int64) (bets []entities.Bet, err error) {
	rows, err := repo.conn.QueryContext(ctx, `
		SELECT
			*
		FROM
			bet
		WHERE
			bet.gambler_id = ?
		`,
		gamblerID,
	)

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
			&bet.RaffleID,
			&bet.BetType,
			&bet.BetModifier,
			&bet.BetPrice,
			&bet.BetCombination,
			&bet.CreatedAt,
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
			bet
		WHERE
			bet.id = ?
		`,
		id,
	)

	err = row.Scan(
		&bet.ID,
		&bet.GamblerID,
		&bet.RaffleID,
		&bet.BetType,
		&bet.BetModifier,
		&bet.BetPrice,
		&bet.BetCombination,
		&bet.CreatedAt,
	)

	return bet, err
}

func (repo *BetRepository) Update(ctx context.Context, bet entities.Bet) (rowsAffected int64, err error) {
	result, err := repo.conn.ExecContext(ctx, `
		UPDATE
			bet
		SET
			bet_type = ?,
			bet_modifier = ?,
			bet_price = ?,
			bet_combination = ?
		WHERE
			id = ?
		`,
		bet.BetType, bet.BetPrice, bet.BetCombination, bet.ID,
	)

	if err != nil {
		return rowsAffected, err
	}

	return result.RowsAffected()
}
