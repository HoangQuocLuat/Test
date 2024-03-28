package router

import (
	handler "thuchanh_go/handler/account"
	chathandler "thuchanh_go/handler/chat"

	"github.com/gin-gonic/gin"
)

type API struct {
	Gin         *gin.Engine
	AccHandler  handler.AccountHandler
	WebHandler  chathandler.Handler
}

func (api *API) SetupRoute() {
	api.Gin.POST("/user/register", api.AccHandler.RegisHandler())
	api.Gin.GET("/user/login", api.AccHandler.LoginHandler())

	api.Gin.POST("/ws/createRoom", api.WebHandler.CreateRoom)
	api.Gin.GET("ws/joinRoom/:roomId", api.WebHandler.JoinRoom)
	api.Gin.GET("ws/getRooms", api.WebHandler.GetRooms)
	api.Gin.GET("ws/getClient/:roomId", api.WebHandler.GetClient)
}
