ALTER TABLE users ADD reset_token VARCHAR(255),
ADD reset_token_expired_at DATETIME;