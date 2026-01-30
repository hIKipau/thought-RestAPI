CREATE TABLE thoughts (
    id BIGINT PRIMARY KEY,
    text TEXT NOT NULL,
    author TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);