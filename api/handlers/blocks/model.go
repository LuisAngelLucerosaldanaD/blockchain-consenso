package blocks

import "time"

type ResAllBlocks struct {
	Error bool     `json:"error"`
	Data  []*Block `json:"data"`
	Code  int      `json:"code"`
	Type  int      `json:"type"`
	Msg   string   `json:"msg"`
}

type Block struct {
	Id                 int64     `json:"id,omitempty"`
	Data               string    `json:"data,omitempty"`
	Nonce              int64     `json:"nonce,omitempty"`
	Difficulty         int32     `json:"difficulty,omitempty"`
	MinedBy            string    `json:"mined_by,omitempty"`
	MinedAt            time.Time `json:"mined_at,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
	Hash               string    `json:"hash,omitempty"`
	PrevHash           string    `json:"prev_hash,omitempty"`
	StatusId           int32     `json:"status_id,omitempty"`
	IdUser             string    `json:"id_user,omitempty"`
	LastValidationDate time.Time `json:"last_validation_date,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}

type resCurrentLottery struct {
	Error bool    `json:"error"`
	Data  Lottery `json:"data"`
	Code  int     `json:"code"`
	Type  int     `json:"type"`
	Msg   string  `json:"msg"`
}

type Lottery struct {
	ID                    string     `json:"id" db:"id" valid:"required,uuid"`
	BlockId               int64      `json:"block_id" db:"block_id" valid:"required"`
	RegistrationStartDate time.Time  `json:"registration_start_date" db:"registration_start_date" valid:"required"`
	RegistrationEndDate   *time.Time `json:"registration_end_date" db:"registration_end_date" valid:"required"`
	LotteryStartDate      *time.Time `json:"lottery_start_date" db:"lottery_start_date" valid:"required"`
	LotteryEndDate        *time.Time `json:"lottery_end_date" db:"lottery_end_date" valid:"required"`
	ProcessEndDate        *time.Time `json:"process_end_date" db:"process_end_date" valid:"required"`
	ProcessStatus         int        `json:"process_status" db:"process_status" valid:"required"`
	CreatedAt             time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at" db:"updated_at"`
}
