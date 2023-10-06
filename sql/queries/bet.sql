-- name: CreateBet :one
INSERT INTO
    bets (gambler_id, bet_type, bet_price, bet_choice)
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: GetBet :one
SELECT
    sqlc.embed(bets),
    sqlc.embed(gamblers)
FROM
    bets
    JOIN gamblers ON bets.gambler_id = gamblers.id
WHERE
    bets.id = $1;

-- name: GetBets :many
SELECT
    sqlc.embed(bets),
    sqlc.embed(gamblers)
FROM
    bets
    JOIN gamblers ON bets.gambler_id = gamblers.id;