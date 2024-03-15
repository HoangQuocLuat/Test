package handler

import (
	"log"
	"net/http"
	"strconv"
	"thuchanh_go/banana"
	"thuchanh_go/logic"
	"thuchanh_go/models"
	"thuchanh_go/redis"
	"thuchanh_go/types/req"
	"thuchanh_go/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AccountHandler struct {
	UserLogic logic.AccLogic
	rd 	      *redis.Redis
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
		userId, _ := uuid.NewUUID()
		user := models.Account{
			ID:       userId.String(),
			Name:     req.Name,
			Phone:    req.Phone,
			Email:    req.Email,
			Username: req.Username,
			Password: hashpassword,
			Token:    "",
		}

		res, err := a.UserLogic.Insert(&gin.Context{}, user)
		if err != nil {
			ctx.JSON(http.StatusConflict, models.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
			return
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
		var req req.UserLoginReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		user, err := a.UserLogic.Select(&gin.Context{}, req)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Sai mật khẩu",
				Data:       nil,
			})
			return
		}

		userID, _ := strconv.ParseInt(user.ID, 10, 64)
		token, err := utils.GenerateJWT(userID, banana.Secretkey)

		user = models.Account{
			Name:  user.Name,
			Phone: user.Phone,
			Email: user.Email,
			Token: token,
		}

		// exists, err := a.rd.Client.Exists(ctx,"user:"+userID+":token").Result()

		ctx.JSON(http.StatusOK, models.Response{
			StatusCode: http.StatusOK,
			Message:    "Đăng nhập thành công",
			Data:       user,
		})
	}
}
