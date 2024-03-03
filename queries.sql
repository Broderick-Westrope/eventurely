-- name: CreateEvent :exec
INSERT INTO events (owner_id, title, description, starts_at, ends_at, location, unique_link, privacy_setting,
                    created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetEvent :one
SELECT *
FROM events
WHERE id = $1
LIMIT 1;

-- name: ListUpcomingOwnedEvents :many
SELECT id,
       owner_id,
       title,
       description,
       starts_at,
       ends_at,
       location,
       unique_link,
       privacy_setting,
       created_at,
       updated_at
FROM events
WHERE starts_at > $1
  AND owner_id = $2
ORDER BY starts_at;

-- name: ListUpcomingInvitedEvents :many
SELECT e.id,
       e.owner_id,
       e.title,
       e.description,
       e.starts_at,
       e.ends_at,
       e.location,
       e.unique_link,
       e.privacy_setting,
       e.created_at,
       e.updated_at,
       i.status
FROM events e
         JOIN invitations i ON e.id = i.event_id
WHERE e.starts_at > $1
  AND i.user_id = $2
ORDER BY starts_at;

-- name: ListPastEvents :many
SELECT e.id,
       e.owner_id,
       e.title,
       e.description,
       e.starts_at,
       e.ends_at,
       e.location,
       e.unique_link,
       e.privacy_setting,
       e.created_at,
       e.updated_at
FROM events e
         JOIN invitations i ON e.id = i.event_id
WHERE e.ends_at < $1
  AND i.user_id = $2
UNION
SELECT e.id,
       e.owner_id,
       e.title,
       e.description,
       e.starts_at,
       e.ends_at,
       e.location,
       e.unique_link,
       e.privacy_setting,
       e.created_at,
       e.updated_at
FROM events e
WHERE e.ends_at < $1
  AND e.owner_id = $2
ORDER BY starts_at DESC;

-- name: UpdateEvent :exec
UPDATE events
set title           = $2,
    description     = $3,
    starts_at       = $4,
    ends_at         = $5,
    location        = $6,
    unique_link     = $7,
    privacy_setting = $8,
    updated_at      = $9
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE
FROM events
WHERE id = $1
RETURNING *;

-- name: UpdateInvitationResponseStatus :exec
UPDATE invitations
SET status = $2, responded_at = $3
WHERE id = $1;