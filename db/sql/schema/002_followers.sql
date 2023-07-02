-- +goose Up
CREATE TABLE user_followers (
    follower_id BIGSERIAL NOT NULL,
    following_id BIGSERIAL NOT NULL,
    PRIMARY KEY (following_id, follower_id),
    FOREIGN KEY (follower_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (following_id) REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE user_followers;