package router

import (
	"github.com/gin-gonic/gin"
	"goGinDemo/controller/homePage"
	"goGinDemo/controller/test"
)

var R *gin.Engine = gin.Default()

func init() {
	R.GET("/ping", testController.Ping)
	R.GET("/helloWorld", homePage.HelloWorld)
}
