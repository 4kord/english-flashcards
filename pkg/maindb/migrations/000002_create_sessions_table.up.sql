CREATE TABLE IF NOT EXISTS sessions (
    id SERIAL PRIMARY KEY,
    refresh_token VARCHAR(255) NOT NULL,
    user_agent TEXT NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id),
    expires_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
);

CREATE UNIQUE INDEX IF NOT EXISTS sessions_session ON sessions (refresh_token);