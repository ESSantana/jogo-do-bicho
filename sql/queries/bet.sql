-- name: CreateBet :exec
INSERT INTO
    bets (gambler_id, bet_type, bet_price, bet_choice)
VALUES
    (?, ?, ?, ?);

-- name: GetBet :one
SELECT
    sqlc.embed(bets),
    sqlc.embed(gamblers)
FROM
    bets
    JOIN gamblers ON bets.gambler_id = gamblers.id
WHERE
    bets.id = ?
    AND bets.deleted_at IS NOT NULL;

-- name: GetBets :many
SELECT
    sqlc.embed(bets),
    sqlc.embed(gamblers)
FROM
    bets
    JOIN gamblers ON bets.gambler_id = gamblers.id
WHERE
    bets.deleted_at IS NOT NULL;

-- name: UpdateBet :exec
UPDATE
    bets
SET
    bet_type = ?,
    bet_price = ?,
    bet_choice = ?
WHERE
    id = ?
    AND bets.deleted_at IS NOT NULL;

-- name: DeleteBet :exec
UPDATE
    bets
SET
    deleted_at = ?
WHERE
    id = ?;