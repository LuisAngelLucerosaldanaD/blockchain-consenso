
-- +migrate Up
CREATE TABLE IF NOT EXISTS bk.validation_table(
    id uuid NOT NULL PRIMARY KEY,
    lottery_id UUID  NOT NULL,
    wallet_id UUID  NOT NULL,
    participants_id UUID  NOT NULL,
    vote BOOLEAN  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS bk.validation_table;
