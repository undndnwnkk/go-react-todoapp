-- +goose Up

CREATE TABLE IF NOT EXISTS categories (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL ,
    name VARCHAR(100) NOT NULL,
    color VARCHAR(10) NOT NULL,
    CONSTRAINT fk_categories_user
                                      FOREIGN KEY (user_id) REFERENCES users(id)
                                      ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS categories;
