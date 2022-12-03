-- name: GetDecks :many
SELECT * FROM decks
WHERE user_id = $1;

-- name: GetPremadeDecks :many
SELECT * FROM decks
WHERE is_premade = true;

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

-- name: DeckAmountUp :exec
UPDATE decks
SET amount = amount + 1
WHERE id = $1;

-- name: DeckAmountDownByCard :exec
UPDATE decks
SET amount = amount - 1
WHERE id = (SELECT deck_id FROM cards WHERE cards.id = $1);