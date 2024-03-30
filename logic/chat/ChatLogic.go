package chatlogic

import (
	"context"
	"thuchanh_go/database"
	"thuchanh_go/logic"
	"thuchanh_go/types/req"

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
