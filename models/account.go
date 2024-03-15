package models

type Account struct {
	ID       string `db:"id" json:"-"`
	Name     string `db:"name" json:"Name"`
	Phone    string `db:"phone" json:"Phone"`
	Email    string `db:"email" json:"Email"`
	Username string `db:"username" json:"Username"`
	Password string `db:"password" json:"-"`
	Token    string `json:"token"`
}


