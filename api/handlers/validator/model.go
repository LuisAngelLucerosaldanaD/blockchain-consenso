package validator

type rqRegisterVote struct {
	WalletID string `json:"wallet_id"`
	Hash     string `json:"hash"`
}

type resRegisterVote struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}
