CREATE TABLE IF NOT EXISTS word_page (
  word_id UUID NOT NULL,
  page_id UUID NOT NULL,
  word_count INT NOT NULL,
  PRIMARY KEY (word_id, page_id),
  CONSTRAINT fk_word FOREIGN KEY (word_id) REFERENCES words(id) ON DELETE CASCADE,
  CONSTRAINT fk_page FOREIGN KEY (page_id) REFERENCES pages(id) ON DELETE CASCADE
);
