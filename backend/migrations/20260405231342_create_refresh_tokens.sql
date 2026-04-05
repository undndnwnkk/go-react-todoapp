-- +goose Up
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS refresh_tokens (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    token_hash VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_refresh_tokens_users
                                          FOREIGN KEY (user_id) REFERENCES users(id)
                                          ON DELETE CASCADE
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE IF EXISTS refresh_tokens;
