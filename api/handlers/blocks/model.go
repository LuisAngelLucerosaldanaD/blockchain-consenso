package blocks

type resAllBlocks struct {
	Error bool     `json:"error"`
	Data  []*Block `json:"data"`
	Code  int      `json:"code"`
	Type  int      `json:"type"`
	Msg   string   `json:"msg"`
}

type Block struct {
	Id                 int64  `json:"id,omitempty"`
	Data               string `json:"data,omitempty"`
	Nonce              int64  `json:"nonce,omitempty"`
	Difficulty         int32  `json:"difficulty,omitempty"`
	MinedBy            string `json:"mined_by,omitempty"`
	MinedAt            string `json:"mined_at,omitempty"`
	Timestamp          string `json:"timestamp,omitempty"`
	Hash               string `json:"hash,omitempty"`
	PrevHash           string `json:"prev_hash,omitempty"`
	StatusId           int32  `json:"status_id,omitempty"`
	IdUser             string `json:"id_user,omitempty"`
	LastValidationDate string `json:"last_validation_date,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`
	UpdatedAt          string `json:"updated_at,omitempty"`
}
