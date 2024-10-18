-- name: CreateGambler :one
INSERT INTO
    gamblers (
        gambler_name,
        document,
        document_type,
        birth_date
    )
VALUES
    (?, ?, ?, ?);

-- name: GetGambler :many
SELECT
    sqlc.embed(gamblers),
    sqlc.embed(bets)
FROM
    gamblers
    LEFT JOIN bets ON gamblers.id = bets.gambler_id
WHERE
    gamblers.id = ?;

-- name: GetGamblers :many
SELECT
    sqlc.embed(gamblers)
FROM
    gamblers;
     
-- name: UpdateGambler :one
UPDATE 
    gamblers
SET 
    gambler_name = ?,
    document = ?,
    document_type = ?,
    birth_date = ?
WHERE 
    id = ? AND gamblers.deleted_at IS NOT NULL;

-- name: DeleteGambler :one
UPDATE 
    gamblers
SET 
    deleted_at = ?
WHERE 
    id = ?;