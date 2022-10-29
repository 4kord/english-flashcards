-- name: GetCards :many
SELECT * FROM cards
WHERE deck_id = $1;

-- name: GetCard :one
SELECT * FROM cards
WHERE id = $1;

-- name: CreateCard :one
INSERT INTO cards (deck_id, english, russian, association, example, transcription, image, image_url, audio, audio_url)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: EditCard :one
UPDATE cards
SET english = $2, russian = $3, association = $4, example = $5, transcription = $6, image = $7, image_url = $8, audio = $9, audio_url = $10
WHERE id = $1
RETURNING *;

-- name: DeleteCard :exec
DELETE FROM cards
WHERE id = $1;

-- name: CopyCard :exec
INSERT INTO cards (deck_id, english, russian, association, example, transcription, image, image_url, audio, audio_url, created_at)
SELECT $2, english, russian, association, example, transcription, image, image_url, audio, audio_url, created_at
FROM cards
WHERE cards.id = $1;