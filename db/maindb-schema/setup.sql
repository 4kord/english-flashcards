CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	admin BOOLEAN DEFAULT(false) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
);

CREATE UNIQUE INDEX IF NOT EXISTS users_email ON users (email);

CREATE TABLE IF NOT EXISTS sessions (
    id SERIAL PRIMARY KEY,
    refresh_token VARCHAR(255) NOT NULL,
    user_agent VARCHAR(15) NOT NULL,
    client_ip VARCHAR(15) NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id),
    expires_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
);

CREATE UNIQUE INDEX IF NOT EXISTS sessions_session ON sessions (session);

CREATE TABLE IF NOT EXISTS decks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    amount INTEGER NOT NULL DEFAULT(0),
    is_premade BOOLEAN NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
);

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