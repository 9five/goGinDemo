package account

import (
	"fmt"
	"goGinDemo/config"
	"goGinDemo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var requestVal struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&requestVal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}

	user, err := model.UR.GetUserByUsername(requestVal.Username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user %s not found", requestVal.Username),
		})
		return
	}

	if user.Password != requestVal.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect password",
		})
		return
	}

	token, err := config.GenerateToken(*user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetUser(ctx *gin.Context) {
	id, ok := ctx.Get("id")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	// username, ok := ctx.Get("username")
	// if !ok {
	//     ctx.JSON(http.StatusUnauthorized, gin.H{})
	//     return
	// }
	user, err := model.UR.GetUserByID(id.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
