package models

type Account struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Phone    string `db:"phone"`
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
}


