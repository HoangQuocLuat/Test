package logic

import (
	"context"
	"thuchanh_go/types/req"
)

type ChatLogic interface {
	Insert(ctx context.Context, room req.CreateRoomReq) (req.CreateRoomReq, error)
}
