CREATE TABLE IF NOT EXISTS progresses (
    id SERIAL PRIMARY KEY,
    card_id INTEGER NOT NULL REFERENCES cards(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    learned BOOLEAN NOT NULL DEFAULT(false),
    selected BOOLEAN NOT NULL DEFAULT(false),
    correct INTEGER NOT NULL DEFAULT(0),
    incorrect INTEGER NOT NULL DEFAULT(0),
    last_viewed_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
);

CREATE UNIQUE INDEX IF NOT EXISTS progresses_card_id_user_id ON progresses (card_id, user_id);