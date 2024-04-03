CREATE TABLE IF NOT EXISTS page_referral (
  referral_id UUID NOT NULL,
  page_id UUID NOT NULL,
  PRIMARY KEY (referral_id, page_id),
  CONSTRAINT fk_referral FOREIGN KEY (referral_id) REFERENCES pages(id) ON DELETE CASCADE,
  CONSTRAINT fk_page FOREIGN KEY (page_id) REFERENCES pages(id) ON DELETE CASCADE
);
