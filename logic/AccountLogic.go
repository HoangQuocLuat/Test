package logic

import (
	"context"
	"thuchanh_go/models"
)

type AccLogic interface {
	Insert(ctx context.Context, user models.Account) (models.Account, error)
}
