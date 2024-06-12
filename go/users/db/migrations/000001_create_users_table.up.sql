CREATE TABLE IF NOT EXISTS users (
    "id" UUID PRIMARY KEY,
    "username" VARCHAR(14) UNIQUE NOT NULL CHECK (LENGTH(username) BETWEEN 4 AND 14),
    "display_name" VARCHAR(20),
    "bio" VARCHAR(160) DEFAULT ''::VARCHAR(160),
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);