-- name: Search :many
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
SELECT url, title, meta_data, total_pages FROM ranked_pages;
