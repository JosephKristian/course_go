package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ipaymupreviews/golang-gin-poc/controller"
)

func AuthRoutes(router *gin.Engine) {
	// Menggunakan /db/v1 sebagai prefix untuk semua rute API
	api := router.Group("/db/v1")
	{
		// Grup rute untuk otentikasi
		auth := api.Group("/auth")
		{
			auth.POST("/login", controller.LoginController)

			auth := api.Group("/register")
			{
				auth.POST("/", controller.RegisterController)
				auth.POST("/email-confirm", controller.EmailConfirmController)
				auth.POST("/resend-email-confirm", controller.ResendEmailConfirmController)
			}
			// Route untuk aktivasi akun pengguna
			auth.POST("/account-activation", controller.AccountActivationController)
			auth.POST("/send-activation-code", controller.ResendEmailConfirmController)

		}
	}
}
