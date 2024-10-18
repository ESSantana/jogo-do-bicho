-- name: GetBetGroup :one
SELECT
    sqlc.embed(bet_groups)
FROM
    bet_groups
WHERE
    bet_groups.id = ?;

-- name: GetBetGroups :many
SELECT
    sqlc.embed(bet_groups)
FROM
    bet_groups;

-- name: GetBetGroupByName :many
SELECT
    sqlc.embed(bet_groups)
FROM
    bet_groups
WHERE
    bet_groups.group_name = ?;