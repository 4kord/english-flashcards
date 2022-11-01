CREATE TABLE IF NOT EXISTS cards (
    id SERIAL PRIMARY KEY,
    deck_id INTEGER NOT NULL REFERENCES decks(id),
    english VARCHAR(255) NOT NULL,
    russian VARCHAR(255) NOT NULL,
    association TEXT,
    example TEXT,
    transcription TEXT,
    image TEXT,
    image_url TEXT,
    audio TEXT,
    audio_url TEXT,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
);