package transactions

type Transaction struct {
	Id        string  `json:"id"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	TypeId    int32   `json:"type_id"`
	Data      string  `json:"data"`
	Block     int64   `json:"block"`
	Files     string  `json:"files"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type ResTransactions struct {
	Error bool           `json:"error"`
	Data  []*Transaction `json:"data"`
	Code  int            `json:"code"`
	Type  int            `json:"type"`
	Msg   string         `json:"msg"`
}

type ResTransaction struct {
	Error bool         `json:"error"`
	Data  *Transaction `json:"data"`
	Code  int          `json:"code"`
	Type  int          `json:"type"`
	Msg   string       `json:"msg"`
}
