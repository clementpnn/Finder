-- name: InsertWord :one
INSERT INTO words (word) VALUES ($1) ON CONFLICT (word) DO NOTHING RETURNING id;

-- name: GetWordId :one
SELECT id FROM words WHERE word = $1;

-- name: InsertWordPage :exec
INSERT INTO word_page (word_id, page_id, word_count) VALUES ($1, $2, $3);