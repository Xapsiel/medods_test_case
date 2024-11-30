package models

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Tokens struct {
	AccessToken  string `json:"access_token,omitempty" `
	RefreshToken string `json:"refresh_token,omitempty"`
	Email        string `json:"email,omitempty"`
}

type RefreshToken struct {
	ID            int    `db:"id"`
	Refresh_token string `db:"refresh_token"`
	Exp           int64  `db:"exp"`
	Ip            string `db:"ip"`
	Email         string `db:"email"`
}
