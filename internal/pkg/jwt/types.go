package jwt

type User struct {
	Username string
	UserId   string
}

type Payload struct {
	UserId   string
	UserName string
	Exp      int64
}

type TokenResponse struct {
	Token string `json:"token"`
}
