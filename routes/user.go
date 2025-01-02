package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ipaymupreviews/golang-gin-poc/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
