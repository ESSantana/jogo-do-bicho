-- name: CreateGambler :one
INSERT INTO
    gamblers (
        gambler_name,
        document,
        document_type,
        birth_date
    )
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: GetGambler :one
SELECT
    sqlc.embed(gamblers),
    sqlc.embed(bets)
FROM
    gamblers
    JOIN bets ON gamblers.id = bets.gambler_id
WHERE
    gamblers.id = $1;

-- name: GetGamblers :many
SELECT
    sqlc.embed(gamblers),
    sqlc.embed(bets)
FROM
    gamblers
    JOIN bets ON gamblers.id = bets.gambler_id;