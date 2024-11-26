package models

type User struct {
	ID int `json:"ID"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	ID            int    `db:"id"`
	Refresh_token string `db:"refresh_token"`
	Exp           int64  `db:"exp"`
	Ip            string `db:"ip"`
}
