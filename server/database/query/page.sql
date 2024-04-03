-- name: IsExistPage :one
SELECT EXISTS(SELECT 1 FROM pages WHERE url = $1);

-- name: GetPageIdByUrl :one
SELECT id FROM pages WHERE url = $1;

-- name: InsertPage :exec
INSERT INTO pages (url, domain_id) VALUES ($1, $2);

-- name: UpdatePageData :one
UPDATE pages SET title = $2, meta_data = $3 WHERE url = $1 RETURNING id;

-- name: UpdatePageWordCount :exec
UPDATE pages SET word_count = $2, is_indexed = TRUE WHERE url = $1;

-- name: InsertPageReferral :exec
INSERT INTO page_referral (page_id, referral_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: GetPageByDomainId :many
SELECT url, id FROM pages WHERE domain_id = $1 ORDER BY url LIMIT 20 OFFSET (($2 - 1) * 20);