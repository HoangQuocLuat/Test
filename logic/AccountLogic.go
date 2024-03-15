package logic

import (
	"context"
	"thuchanh_go/models"
	"thuchanh_go/types/req"
)

type AccLogic interface {
	Insert(ctx context.Context, user models.Account) (models.Account, error)
	Select(ctx context.Context, req req.UserLoginReq) (models.Account, error)
}
