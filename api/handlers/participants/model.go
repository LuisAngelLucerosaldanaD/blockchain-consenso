package participants

type responseRegisterParticipant struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}

type requestRegisterParticipant struct {
	Amount int64 `json:"amount"`
}
