package model

type (
	UserInfo struct {
		UserName   string `json:"username"`
		PasswdHash []byte `json:"password_hash"`
		Salt       []byte `json:"salt"`
		Token      string `json:"token"`
	}
)
