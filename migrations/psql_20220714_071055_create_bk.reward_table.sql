
-- +migrate Up
CREATE TABLE IF NOT EXISTS bc.reward_table(
    id uuid NOT NULL PRIMARY KEY,
    lottery_id UUID  NOT NULL,
    id_wallet UUID  NOT NULL,
    amount BIGINT  NOT NULL,
    block_id BIGINT  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

ALTER TABLE bc.reward_table ADD CONSTRAINT fk_reward_table_lottery FOREIGN KEY (lottery_id) REFERENCES bc.lottery_table(id);
ALTER TABLE bc.reward_table ADD CONSTRAINT fk_reward_table_wallets FOREIGN KEY (id_wallet) REFERENCES auth.wallets(id);

-- +migrate Down
DROP TABLE IF EXISTS bk.reward_table;
