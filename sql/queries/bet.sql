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
    bets.id = $1
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

-- name: UpdateBet :one
UPDATE 
    bets
SET 
    bet_type = $1,
    bet_price = $2,
    bet_choice = $3
WHERE 
    id = $4 AND bets.deleted_at IS NOT NULL RETURNING *;

-- name: DeleteBet :one
UPDATE 
    bets
SET 
    deleted_at = $1
WHERE 
    id = $2 RETURNING *;