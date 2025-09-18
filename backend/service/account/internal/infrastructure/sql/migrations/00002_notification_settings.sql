-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS account.notification_settings(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    send_email BOOLEAN,
    send_push_notifications BOOLEAN,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    is_archived BOOLEAN
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS account.notification_settings;
-- +goose StatementEnd
