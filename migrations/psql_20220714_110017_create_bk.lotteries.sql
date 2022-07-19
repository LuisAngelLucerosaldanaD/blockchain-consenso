
-- +migrate Up
CREATE TABLE IF NOT EXISTS bc.lotteries(
    id uuid NOT NULL PRIMARY KEY,
    block_id BIGINT  NOT NULL,
    registration_start_date TIMESTAMP  NOT NULL,
    registration_end_date TIMESTAMP  ,
    lottery_start_date TIMESTAMP  ,
    lottery_end_date TIMESTAMP  ,
    process_end_date TIMESTAMP  ,
    process_status INTEGER  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

ALTER TABLE bc.lotteries ADD CONSTRAINT fk_lotteries_process FOREIGN KEY (process_status) REFERENCES cfg.dictionaries(id);

-- +migrate Down
DROP TABLE IF EXISTS bc.lottery_table;
