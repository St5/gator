-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    published_at TIMESTAMP,
    feed_id UUID NOT NULL,
    FOREIGN KEY (feed_id) REFERENCES feeds(id)
);

-- +goose Down
DROP TABLE posts;