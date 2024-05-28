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

-- name: GetGambler :many
SELECT
    sqlc.embed(gamblers),
    sqlc.embed(bets)
FROM
    gamblers
    LEFT JOIN bets ON gamblers.id = bets.gambler_id
WHERE
    gamblers.id = $1;

-- name: GetGamblers :many
SELECT
    sqlc.embed(gamblers)
FROM
    gamblers;
    
-- name: UpdateGambler :one
UPDATE 
    gamblers
SET 
    gambler_name = $1,
    document = $2,
    document_type = $3,
    birth_date = $4
WHERE 
    id = $5 AND gamblers.deleted_at IS NOT NULL RETURNING *;

-- name: DeleteGambler :one
UPDATE 
    gamblers
SET 
    deleted_at = $1
WHERE 
    id = $2 RETURNING *;