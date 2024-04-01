package models



type Message struct {
	ID      string `db:"id"`
	RoomID  string `db:"roomid"`
	Sender  string `db:"sender"`
	Content string `db:"content"`
	
}
