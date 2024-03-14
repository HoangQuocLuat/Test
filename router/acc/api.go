package router

import (
	handler "thuchanh_go/handler/account"

	"github.com/gin-gonic/gin"
)

type API struct {
	Gin *gin.Engine
	AccHandler handler.AccountHandler
}
func (api *API) SetupRoute() {
	api.Gin.POST("/user/register", api.AccHandler.RegisHandler())
	api.Gin.POST("/user/login", api.AccHandler.RegisHandler())
}
