-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS account.contact_information(
    id UUID PRIMARY KEY NOT NULL,
    profile_id UUID,
    contact_information_type VARCHAR(255),
    is_public BOOLEAN,
    contact_information VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    is_archived BOOLEAN,
    CONSTRAINT fk_profile
        FOREIGN KEY(profile_id)
        REFERENCES account.profile(id)
        ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS account.contact_information;
-- +goose StatementEnd
