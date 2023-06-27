
-- +migrate Up
CREATE TABLE IF NOT EXISTS cfg.blion_access(
    id uuid NOT NULL PRIMARY KEY,
    key VARCHAR (50) NOT NULL,
    ttl VARCHAR (10)  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS cfg.blion_access;
