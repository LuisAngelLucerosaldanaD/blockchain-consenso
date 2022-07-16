
-- +migrate Up
CREATE TABLE IF NOT EXISTS auth.frozen_money(
    id uuid NOT NULL PRIMARY KEY,
    wallet_id UUID  NOT NULL,
    amount BIGINT  NOT NULL,
    lottery_id UUID  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS auth.frozen_money;
