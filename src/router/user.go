package router

import (
	"github.com/CountryMarket/CountryMarket-backend/controller"
	"github.com/CountryMarket/CountryMarket-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.RouterGroup) {
	g.GET("/code", controller.UserLogin)
	g.GET("/validate", middleware.JWTMiddleware(), controller.UserValidate)
	g.GET("/profile", middleware.JWTMiddleware(), controller.UserGetProfile)
}
