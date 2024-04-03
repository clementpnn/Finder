-- name: IsExistDomain :one
SELECT EXISTS(SELECT 1 FROM domains WHERE domain = $1);

-- name: InsertDomain :one
INSERT INTO domains (domain) VALUES ($1) RETURNING id;

-- name: DeleteDomain :exec
DELETE FROM domains WHERE domain = $1;