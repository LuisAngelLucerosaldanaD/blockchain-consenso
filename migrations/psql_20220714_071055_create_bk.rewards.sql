
-- +migrate Up
CREATE TABLE IF NOT EXISTS bc.rewards(
    id uuid NOT NULL PRIMARY KEY,
    lottery_id UUID  NOT NULL,
    id_wallet UUID  NOT NULL,
    amount float8  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

ALTER TABLE bc.rewards ADD CONSTRAINT fk_rewards_lottery FOREIGN KEY (lottery_id) REFERENCES bc.lotteries(id);
ALTER TABLE bc.rewards ADD CONSTRAINT fk_rewards_wallets FOREIGN KEY (id_wallet) REFERENCES auth.wallets(id);

-- +migrate Down
DROP TABLE IF EXISTS bc.rewards;
