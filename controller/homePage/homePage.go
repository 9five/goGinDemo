package homePage

import (
	// "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context) {
	data := map[string]string{"msg": "Hello World"}
	ctx.JSON(http.StatusOK, data)
}
