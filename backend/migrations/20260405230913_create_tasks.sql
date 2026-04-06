-- +goose Up

CREATE TABLE IF NOT EXISTS tasks (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    category_id uuid,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    status VARCHAR(100) NOT NULL,
    priority INT NOT NULL,
    due_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_tasks_users
                                 FOREIGN KEY (user_id) REFERENCES users(id)
                                 ON DELETE CASCADE,
    CONSTRAINT fk_tasks_categories
                                 FOREIGN KEY (category_id) REFERENCES categories(id)
                                 ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS tasks;
