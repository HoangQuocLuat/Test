package handler

import (
	"log"
	"net/http"
	"thuchanh_go/logic"
	"thuchanh_go/models"
	"thuchanh_go/types/req"
	"thuchanh_go/utils"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	UserLogic logic.AccLogic
}

func (a *AccountHandler) RegisHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req req.UserRegisReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		hashpassword, err := utils.HashPassword(req.Password)
		if err != nil {
			log.Fatal(err)
		}

		user := models.Account{
			Name:     req.Name,
			Phone:    req.Phone,
			Email:    req.Email,
			Username: req.Username,
			Password: hashpassword,
		}

		res, err := a.UserLogic.Insert(&gin.Context{}, user)
		if err != nil {
			ctx.JSON(http.StatusConflict, models.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		ctx.JSON(http.StatusOK, models.Response{
			StatusCode: http.StatusOK,
			Message:    "Đăng ký thành công",
			Data:       res,
		})
	}
}

func (a *AccountHandler) LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
