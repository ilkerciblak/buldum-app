-- name: CreateProfile :exec
INSERT INTO account.profile (id, user_name, avatar_url, created_at, is_archived)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
);



-- CREATE TABLE IF NOT EXISTS account.profile(
--     id UUID PRIMARY KEY,
--     user_id UUID NOT NULL,
--     user_name VARCHAR(255) NOT NULL UNIQUE,
--     avatar_url TEXT,
--     created_at TIMESTAMPTZ NOT NULL,
--     updated_at TIMESTAMPTZ,
--     deleted_at TIMESTAMPTZ,
--     is_archived BOOLEAN
-- );
-- -- +goose StatementEnd

