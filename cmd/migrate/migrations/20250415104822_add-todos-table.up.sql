CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);