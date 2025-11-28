-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS teams (
     id          SERIAL PRIMARY KEY,
     name        TEXT NOT NULL UNIQUE,
     created_at  TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users (
     id          TEXT PRIMARY KEY,
     username    TEXT NOT NULL,
     is_active   BOOLEAN NOT NULL DEFAULT true,
     created_at  TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS team_members (
    team_id     INTEGER NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    user_id     TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (team_id, user_id)
);

CREATE TABLE IF NOT EXISTS pull_requests (
    id            TEXT PRIMARY KEY,
    name          TEXT NOT NULL,
    author_id     TEXT NOT NULL REFERENCES users(id),
    status        TEXT NOT NULL CHECK (status IN ('OPEN','MERGED')),
    created_at    TIMESTAMPTZ DEFAULT now(),
    merged_at     TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS pull_request_reviewers (
    pull_request_id TEXT NOT NULL REFERENCES pull_requests(id) ON DELETE CASCADE,
    user_id         TEXT NOT NULL REFERENCES users(id),
    assigned_at     TIMESTAMPTZ DEFAULT now(),
    PRIMARY KEY (pull_request_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_team_members_user_id
    ON team_members(user_id);

CREATE INDEX IF NOT EXISTS idx_pull_request_reviewers_user_id
    ON pull_request_reviewers(user_id);

CREATE INDEX IF NOT EXISTS idx_pull_requests_author_id
    ON pull_requests(author_id);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS pull_request_reviewers;
DROP TABLE IF EXISTS pull_requests;
DROP TABLE IF EXISTS team_members;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS teams;

-- +goose StatementEnd
