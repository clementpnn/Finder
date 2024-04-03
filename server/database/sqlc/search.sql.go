// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: search.sql

package sqlc

import (
	"context"
	"encoding/json"

	"github.com/lib/pq"
)

const search = `-- name: Search :many
WITH input_words AS (
    SELECT id FROM words WHERE word IN (SELECT unnest($1::text[]))
),
word_in_pages AS (
    SELECT 
        wp.page_id, 
        w.id AS word_id,
        wp.word_count AS word_frequency,
        p.word_count AS total_words_on_page
    FROM word_page wp
    INNER JOIN input_words w ON wp.word_id = w.id
    INNER JOIN pages p ON wp.page_id = p.id
),
total_pages AS (
    SELECT COUNT(*)::float AS count FROM pages
),
word_in_pages_count AS (
    SELECT 
        word_id, 
        COUNT(DISTINCT page_id)::float AS count
    FROM word_in_pages
    GROUP BY word_id
),
tf AS (
    SELECT 
        page_id, 
        word_id,
        (word_frequency / total_words_on_page)::float AS tf
    FROM word_in_pages
),
idf AS (
    SELECT 
        word_id,
        LOG((SELECT count FROM total_pages) / MAX(count)) AS idf
    FROM word_in_pages_count
    GROUP BY word_id
),
tfidf AS (
    SELECT 
        tf.page_id, 
        SUM(tf.tf * idf.idf) AS tfidf
    FROM tf
    JOIN idf ON tf.word_id = idf.word_id
    GROUP BY tf.page_id
),
page_rank AS (
    SELECT 
        pr.page_id, 
        COUNT(pr.referral_id) AS incoming_links_count
    FROM page_referral pr
    GROUP BY pr.page_id
),
ranked_pages AS (
    SELECT 
        p.url, 
        p.title,
        p.meta_data,
        tfidf.tfidf,
        COALESCE(pr.incoming_links_count, 0) AS page_rank,
        CEIL(COUNT(*) OVER() / 20.0) AS total_pages
    FROM tfidf
    INNER JOIN pages p ON tfidf.page_id = p.id
    LEFT JOIN page_rank pr ON p.id = pr.page_id
    ORDER BY tfidf DESC, page_rank DESC
    LIMIT 20 OFFSET $2
)
SELECT url, title, meta_data, total_pages FROM ranked_pages
`

type SearchParams struct {
	Column1 []string
	Offset  int32
}

type SearchRow struct {
	Url        string
	Title      string
	MetaData   json.RawMessage
	TotalPages float64
}

func (q *Queries) Search(ctx context.Context, arg SearchParams) ([]SearchRow, error) {
	rows, err := q.db.QueryContext(ctx, search, pq.Array(arg.Column1), arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchRow
	for rows.Next() {
		var i SearchRow
		if err := rows.Scan(
			&i.Url,
			&i.Title,
			&i.MetaData,
			&i.TotalPages,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}