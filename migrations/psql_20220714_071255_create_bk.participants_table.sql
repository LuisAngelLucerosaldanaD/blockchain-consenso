
-- +migrate Up
CREATE TABLE IF NOT EXISTS bk.participants_table(
    id uuid NOT NULL PRIMARY KEY,
    lottery_id UUID  NOT NULL,
    wallet_id UUID  NOT NULL,
    amount BIGINT  NOT NULL,
    accepted BOOLEAN  NOT NULL,
    type_charge INTEGER  NOT NULL,
    returned BOOLEAN  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS bk.participants_table;
