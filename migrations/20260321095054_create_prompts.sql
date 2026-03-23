-- +goose Up
CREATE TABLE prompts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE,
    prompt TEXT NOT NULL
);

-- +goose Down
SELECT 'down SQL query';
