
-- +migrate Up
CREATE TABLE IF NOT EXISTS bc.participants(
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

ALTER TABLE bc.participants ADD CONSTRAINT fk_participants_lottery FOREIGN KEY (lottery_id) REFERENCES bc.lotteries(id);
ALTER TABLE bc.participants ADD CONSTRAINT fk_participants_type_charge FOREIGN KEY (type_charge) REFERENCES cfg.dictionaries(id);
ALTER TABLE bc.participants ADD CONSTRAINT fk_participants_wallets FOREIGN KEY (wallet_id) REFERENCES auth.wallets(id);

-- +migrate Down
DROP TABLE IF EXISTS bc.participants;
