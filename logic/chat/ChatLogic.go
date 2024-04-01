package chatlogic

import (
	"context"
	"database/sql"
	"thuchanh_go/banana"
	"thuchanh_go/database"
	"thuchanh_go/logic"
	"thuchanh_go/types/req"
	"thuchanh_go/types/res"

	"github.com/labstack/gommon/log"
)

type ChatRoomLogic struct {
	sql *database.Sql
}

func NewChatLogic(sql *database.Sql) logic.ChatLogic {
	return &ChatRoomLogic{
		sql: sql,
	}
}

func (c *ChatRoomLogic) Insert(ctx context.Context, room req.CreateRoomReq) (req.CreateRoomReq, error) {
	statement := `INSERT INTO room (id, name) VALUES(:id, :name)`
	_, err := c.sql.Db.NamedExecContext(ctx, statement, room)
	if err != nil {
		log.Error(err.Error())
	}
	return room, nil
}

func (c *ChatRoomLogic) Select(ctx context.Context, req req.CreateRoomReq) (res.RoomRes, error) {
	room := res.RoomRes{}
	statement := `SELECT * FROM room WHERE name = $1`
	err := c.sql.Db.GetContext(ctx, &room, statement, req.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return room, banana.RoomNotFoud
		}
		log.Error(err.Error())
		return room, err
	}
	return room, nil
}

func (c *ChatRoomLogic) InsertMess(ctx context.Context, mess req.Message) error {
	statement := `INSERT INTO message (roomid, sender, content) 
	            VALUES(:roomid, :sender, :content)`
	_, err := c.sql.Db.NamedExecContext(ctx, statement, mess)
	if err != nil {
		log.Error(err.Error())
	}
	return nil
}
