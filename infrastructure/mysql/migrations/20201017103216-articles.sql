
-- +migrate Up
CREATE TABLE IF NOT EXISTS articles (
    id bigint AUTO_INCREMENT NOT NULL,
    title VARCHAR(255),
    body TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS articles;