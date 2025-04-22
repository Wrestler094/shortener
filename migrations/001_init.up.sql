CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36),
    short_url VARCHAR(20) NOT NULL UNIQUE,
    original_url TEXT NOT NULL UNIQUE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_original_url_unique ON urls (original_url);
