
-- +migrate Up
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL,
    name VARCHAR( 255 ) DEFAULT '' NOT NULL,
    description VARCHAR( 255 ) DEFAULT '' NOT NULL,
    status_id INTEGER DEFAULT 0 NOT NULL,
    user_id INTEGER DEFAULT 0 NOT NULL,
    due_date TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY ( id )
);

CREATE INDEX task_user_id_index ON tasks (user_id);

-- +migrate Down
DROP TABLE tasks;