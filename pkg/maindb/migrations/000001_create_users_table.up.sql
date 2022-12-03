CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	admin BOOLEAN DEFAULT(false) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
);

CREATE UNIQUE INDEX IF NOT EXISTS users_email ON users (email);

INSERT INTO users (email, password, admin) VALUES ('admin@admin.com', '$2a$10$W4PLUzvGWC7.goqkq8ln5OQJTfxl.U.20gS5Kya5L474DvMtwYiOG', true)