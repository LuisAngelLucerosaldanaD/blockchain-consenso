package miner

import "time"

type rqRegisterMined struct {
	Hash     string `json:"hash"`
	Nonce    int64  `json:"nonce"`
	WalletID string `json:"wallet_id"`
}

type responseRegisterMined struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}

type responseGetBlock struct {
	Error bool             `json:"error"`
	Data  *DataBlockToMine `json:"data"`
	Code  int              `json:"code"`
	Type  int              `json:"type"`
	Msg   string           `json:"msg"`
}

type DataBlockToMine struct {
	ID         int64     `json:"id"`
	Data       []byte    `json:"data"`
	Timestamp  time.Time `json:"timestamp"`
	PrevHash   []byte    `json:"prev_hash"`
	Difficulty int       `json:"difficulty"`
	Cuota      float64   `json:"cuota"`
}
