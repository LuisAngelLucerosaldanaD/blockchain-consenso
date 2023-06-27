package sign

type ReqSign struct {
	Key string `json:"key"`
}

type ResSign struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}

type ResExportSign struct {
	Error bool   `json:"error"`
	Data  []Sign `json:"data"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}

type Sign struct {
	Id  string `json:"id"`
	Key string `json:"key"`
}

type ReqExportSign struct {
	Id []string `json:"id"`
}
