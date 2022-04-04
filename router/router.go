package router

import (
	"goGinDemo/controller/account"
	"goGinDemo/controller/homePage"
	"goGinDemo/controller/test"
	"goGinDemo/middleware"

	"github.com/gin-gonic/gin"
)

var R *gin.Engine = gin.Default()

func init() {
	R = accountGroup(R)
	R.GET("/ping", testController.Ping)
	R.GET("/helloWorld", homePage.HelloWorld)
}

func accountGroup(router *gin.Engine) *gin.Engine {
	accountAPI := router.Group("/account")
	accountAPI.POST("/login", account.Login)
	accountAPI.GET("/id", middleware.VerifyToken(), account.GetUser)
	return router
}
