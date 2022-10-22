CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	role VARCHAR(255) DEFAULT("user") NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
);
ALTER TABLE users
ADD CONSTRAINT uc_users_email UNIQUE(email);

CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    session VARCHAR(255) NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id),
    ip VARCHAR(15) NOT NULL,
    expires_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
);
ALTER TABLE sessions
ADD CONSTRAINT uc_sessions_sesssion UNIQUE(session);

CREATE TABLE decks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    amount INTEGER NOT NULL DEFAULT(0),
    is_premade BOOLEAN NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
);
ALTER TABLE decks
ADD CONSTRAINT uc_decks_name UNIQUE(name);

CREATE TABLE cards (
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

CREATE TABLE progresses (
    id SERIAL PRIMARY KEY,
    card_id INTEGER NOT NULL REFERENCES cards(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    learned BOOLEAN NOT NULL DEFAULT(false),
    selected BOOLEAN NOT NULL DEFAULT(false),
    correct INTEGER NOT NULL DEFAULT(0),
    incorrect INTEGER NOT NULL DEFAULT(0),
    last_viewed_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
)