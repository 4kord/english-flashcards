-- name: GetDecks :many
SELECT * FROM decks
WHERE user_id = $1;

-- name: GetPremadeDecks :many
SELECT * FROM decks
WHERE premade = true;

-- name: CreateDeck :one
INSERT INTO decks (user_id, name, is_premade)
VALUES ($1, $2, $3)
RETURNING *;

-- name: EditDeck :one
UPDATE decks
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteDeck :exec
DELETE FROM decks
WHERE id = $1;