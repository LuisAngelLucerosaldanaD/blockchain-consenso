package participants

type responseRegisterParticipant struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}

type requestRegisterParticipant struct {
	Amount float64 `json:"amount"`
}

type resParticipant struct {
	Error bool             `json:"error"`
	Data  *InfoParticipant `json:"data"`
	Code  int              `json:"code"`
	Type  int              `json:"type"`
	Msg   string           `json:"msg"`
}

type InfoParticipant struct {
	Accepted bool `json:"accepted"`
	Charge   int  `json:"charge"`
}
