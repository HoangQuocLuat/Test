package logic

import (
	"context"
	"thuchanh_go/types/req"
	"thuchanh_go/types/res"
)

type ChatLogic interface {
	Insert(ctx context.Context, room req.CreateRoomReq) (req.CreateRoomReq, error)
	Select(ctx context.Context, room req.CreateRoomReq) (res.RoomRes, error)

	InsertMess(ctx context.Context, mess req.Message) (error)
}
