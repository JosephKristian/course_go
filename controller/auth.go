package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/ipaymupreviews/golang-gin-poc/models"
)

func LoginController(c *gin.Context) {
	c.String(200, "Hello PRENS LOGIN")
}

func RegisterController(c *gin.Context) {
	var userInput models.RegisterInput
	c.String(200, "Hello PRENS LOGIN")
	// Ambil Bearer Token dari header Authorization
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(400, gin.H{
			"status":    "error",
			"data":      nil,
			"message":   "Authorization token is required",
			"errorCode": 400,
		})
		return
	}

	// Bind input multipart-form data ke struct
	if err := c.ShouldBind(&userInput); err != nil {
		log.Printf("[ERROR] Invalid input data: %v", err)
		c.JSON(422, gin.H{
			"status":    "error",
			"data":      nil,
			"message":   "Invalid input data",
			"errorCode": 422,
			"errors":    err.Error(),
		})
		return
	}

	log.Println("[INFO] Input data validated.")

	// Panggil service untuk registrasi user
	registeredUser, err := services.registerService

	if err != nil {
		log.Printf("[ERROR] Registration failed: %v", err)
		c.JSON(500, gin.H{
			"status":    "error",
			"data":      nil,
			"message":   "Registration failed",
			"errorCode": 500,
			"errors":    err.Error(),
		})
		return
	}

	log.Printf("[INFO] User registered successfully: %s", registeredUser.Email)

	// Respons sukses setelah berhasil registrasi
	c.Header("Location", "/db/v1/auth/register/"+registeredUser.Email)
	c.JSON(201, gin.H{
		"status":  "success",
		"data":    userInput,
		"message": "User registered successfully",
	})
}

func AccountActivationController(c *gin.Context) {
	c.String(200, "Hello PRENS AccountActivation")
}

func ResendActivationCodeController(c *gin.Context) {
	c.String(200, "Hello PRENS ResendActivationCode")
}
func EmailConfirmController(c *gin.Context) {
	c.String(200, "Hello PRENS ResendActivationCode")
}
func ResendEmailConfirmController(c *gin.Context) {
	c.String(200, "Hello PRENS ResendEmailConfirmController")
}
