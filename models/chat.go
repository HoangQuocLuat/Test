package models

type Room struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

type Message struct {
	ID      string `db:"id"`
	RoomID  string `db:"roomId"`
	UserID  string `db:"userId"`
	Content string `db:"content"`
}
