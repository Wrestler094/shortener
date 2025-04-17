CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    short_url VARCHAR(20) NOT NULL UNIQUE,
    original_url TEXT NOT NULL UNIQUE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_original_url_unique ON urls (original_url);
