-- name: CreateProfile :exec
INSERT INTO account.profile (id, user_name, avatar_url, created_at, is_archived)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
);

-- name: GetProfileById :one
SELECT * FROM account.profile
WHERE id=$1;


-- name: GetAllProfile :many
SELECT * FROM account.profile
ORDER BY $1
LIMIT $2
OFFSET $3;

-- name: ArchiveProfile :exec
UPDATE account.profile
SET is_archived=true
WHERE id=$1;

-- name: DeleteProfile :exec
DELETE FROM account.profile
WHERE id=$1;

-- name: UpdateProfile :exec
UPDATE account.profile
SET 
    user_name=$1,
    avatar_url=$2,
    updated_at=$3
WHERE id=$4;

