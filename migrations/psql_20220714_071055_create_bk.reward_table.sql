
-- +migrate Up
CREATE TABLE IF NOT EXISTS bk.reward_table(
    id uuid NOT NULL PRIMARY KEY,
    lottery_id UUID  NOT NULL,
    id_wallet UUID  NOT NULL,
    amount BIGINT  NOT NULL,
    block_id BIGINT  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS bk.reward_table;
