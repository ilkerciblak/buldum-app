-- name: CreateContactInformation :exec
INSERT INTO account.contact_information (
    id,
    profile_id,
    contact_information_type,
    is_public,
    contact_information,
    created_at,
    is_archived
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
);


-- name: UpdateContactInformation :exec
UPDATE account.contact_information
SET
    is_public=$1,
    contact_information=$2,
    updated_at=$3
WHERE profile_id=$4 AND contact_information_type=$5;

-- name: ArchiveContactInformation :exec
UPDATE account.contact_information
SET
    is_archived=$1,
    deleted_at=$2,
    updated_at=$3
WHERE id=$4;


-- name: GetAllContactInformation :many
SELECT * FROM account.contact_information
WHERE 
    ($1::uuid = '00000000-0000-0000-0000-000000000000' OR profile_id = $1)
    AND ($2::text = '' OR contact_information_type = $2)
    AND ($3::bool = true OR is_public = $3)
    AND ($4::bool = false OR is_archived = $4)
;

-- name: GetByID :one
SELECT * FROM account.contact_information
WHERE id=$1;
