-- name: CreateRefreshToken :one
INSERT INTO refreshTokens (token, created_at, updated_at, expires_at, revoked_at, user_id)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    NULL,
    $3
)
RETURNING *;

-- name: GetRefreshToken :one
SELECT token, created_at, updated_at, expires_at, revoked_at, user_id
FROM refreshTokens
WHERE token = $1;

-- name: RevokeRefreshToken :exec
UPDATE refreshTokens
SET revoked_at = NOW(), updated_at = NOW()
WHERE token = $1;