-- +goose Up
-- +goose StatementBegin

-- creating users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(127) NOT NULL UNIQUE,
    password VARCHAR(127) NOT NULL,
    email VARCHAR(255)
    );

-- creating posts table
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id INTEGER NOT NULL REFERENCES users(id),
    comments_allowed BOOLEAN NOT NULL
    );

-- creating comments table
CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    post_id INTEGER,
    user_id INTEGER,
    text TEXT,
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
    );

-- creating comments_parent_childs_ids table
CREATE TABLE IF NOT EXISTS comments_parent_childs_ids (
    parent_id INTEGER,
    children_id INTEGER,
    level INTEGER,
    PRIMARY KEY (parent_id, children_id),
    FOREIGN KEY (parent_id) REFERENCES comments(id),
    FOREIGN KEY (children_id) REFERENCES comments(id)
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- deleting comments_parent_childs_ids table
DROP TABLE IF EXISTS comments_parent_childs_ids;

-- deleting comments table
DROP TABLE IF EXISTS comments;

-- deleting posts table
DROP TABLE IF EXISTS posts;

-- deleting users table
DROP TABLE IF EXISTS users;

-- +goose StatementEnd