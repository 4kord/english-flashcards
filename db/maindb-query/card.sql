-- name: GetCards :many
SELECT * FROM cards
WHERE deck_id = $1;

-- name: CreateCard :one
INSERT INTO cards (deck_id, english, russian, association, example, transcription, image, image_url, audio, audio_url)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;