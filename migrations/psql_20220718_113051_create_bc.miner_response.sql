
-- +migrate Up
CREATE TABLE IF NOT EXISTS bc.miner_response(
    id uuid NOT NULL PRIMARY KEY,
    lottery_id UUID  NOT NULL,
    participants_id UUID  NOT NULL,
    hash VARCHAR (255) NOT NULL,
    status INTEGER  NOT NULL,
    nonce BIGINT NOT NULL,
    difficulty INTEGER  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

ALTER TABLE bc.miner_response ADD CONSTRAINT fk_miner_response_lottery FOREIGN KEY (lottery_id) REFERENCES bc.lotteries(id);
ALTER TABLE bc.miner_response ADD CONSTRAINT fk_miner_response_participants FOREIGN KEY (participants_id) REFERENCES bc.participants(id);
ALTER TABLE bc.miner_response ADD CONSTRAINT fk_miner_response_process FOREIGN KEY (status) REFERENCES cfg.dictionaries(id);

-- +migrate Down
DROP TABLE IF EXISTS bc.miner_response;
