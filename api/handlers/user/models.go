package user

type rqLogin struct {
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responseLogin struct {
	Error bool   `json:"error"`
	Data  Token  `json:"data"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type resGetWallets struct {
	Error bool      `json:"error"`
	Data  []*Wallet `json:"data"`
	Code  int       `json:"code"`
	Type  int       `json:"type"`
	Msg   string    `json:"msg"`
}

type Wallet struct {
	Id               string `json:"id,omitempty"`
	Mnemonic         string `json:"mnemonic,omitempty"`
	RsaPublic        string `json:"rsa_public,omitempty"`
	RsaPrivate       string `json:"rsa_private,omitempty"`
	RsaPublicDevice  string `json:"rsa_public_device,omitempty"`
	RsaPrivateDevice string `json:"rsa_private_device,omitempty"`
	IpDevice         string `json:"ip_device,omitempty"`
	StatusId         int32  `json:"status_id,omitempty"`
	IdentityNumber   string `json:"identity_number,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
}

type resAccount struct {
	Error bool       `json:"error"`
	Data  Accounting `json:"data"`
	Code  int        `json:"code"`
	Type  int        `json:"type"`
	Msg   string     `json:"msg"`
}

type Accounting struct {
	Id        string  `json:"id,omitempty"`
	IdWallet  string  `json:"id_wallet,omitempty"`
	Amount    float64 `json:"amount,omitempty"`
	IdUser    string  `json:"id_user,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
}

type resFreezeMoney struct {
	Error bool   `json:"error"`
	Data  int64  `json:"data"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}

type responseAnny struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}

type requestActivateUser struct {
	Code string `json:"code"`
}

type ReqChangePwd struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

type ChangePwd struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
