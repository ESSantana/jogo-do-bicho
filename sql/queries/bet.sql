-- name: CreateBet :exec
INSERT INTO
    bets (id, animal, bet_number, bet_price)
VALUES
    ('?', '?', '?', '?');

-- name: GetBet :one
SELECT
    *
FROM
    bets
WHERE
    id = '?';

-- name: GetBets :many
SELECT
    *
FROM
    bets;