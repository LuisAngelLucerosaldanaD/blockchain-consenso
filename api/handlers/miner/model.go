package miner

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
