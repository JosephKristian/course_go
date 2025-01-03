package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/ipaymupreviews/golang-gin-poc/config"
	"gitlab.com/ipaymupreviews/golang-gin-poc/models"
	"gitlab.com/ipaymupreviews/golang-gin-poc/services"
)

func LoginController(c *gin.Context) {
	c.String(200, "Hello PRENS LOGIN")
}

func RegisterController(c *gin.Context) {
	var userInput models.RegisterInput

	// Ambil Bearer Token dari header Authorization
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(400, gin.H{
			"status":     "error",
			"statusCode": 400,
			"message":    "Authorization token is required",
		})
		return
	}

	// Bind input multipart-form data ke struct
	if err := c.ShouldBind(&userInput); err != nil {
		c.JSON(422, gin.H{
			"status":     "error",
			"statusCode": 422,
			"errors":     err.Error(),
			"message":    "Invalid input data",
		})
		return
	}

	// Panggil service untuk registrasi user
	registeredUser, err := services.RegisterService(&userInput)
	if err != nil {
		c.JSON(500, gin.H{
			"status":     "error",
			"statusCode": 500,
			"message":    err.Error(),
		})
		return
	}

	// Respons sukses setelah berhasil registrasi
	c.Header("Location", "/db/v1/auth/register/"+registeredUser.Email)
	c.JSON(201, gin.H{
		"status":     "success",
		"statusCode": 201,
		"data":       registeredUser,
		"message":    "User registered successfully",
	})
}

func AccountActivationController(c *gin.Context) {

	var userInput models.AccountActivation

	// Ambil Bearer Token dari header Authorization
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(400, gin.H{
			"status":     "error",
			"statusCode": 400,
			"message":    "Authorization token is required",
		})
		return
	}

	// Bind input multipart-form data ke struct
	if err := c.ShouldBind(&userInput); err != nil {
		c.JSON(422, gin.H{
			"status":     "error",
			"statusCode": 422,
			"message":    "Invalid input data",
			"errors":     err.Error(),
		})
		return
	}

	isVerified, err := services.VerifyOtp(userInput.EmailOrPhone, userInput.VerificationCode)

	if err != nil {
		log.Printf("[ERROR] OTP verification failed for %s: %v", userInput.EmailOrPhone, err)
		c.JSON(400, gin.H{"status": "error", "message": "OTP verification failed", "error": err.Error()})
		return
	}

	if isVerified {
		// Update kolom is_active di tabel users
		result := config.DB.Model(&models.User{}).
			Where("email = ? OR phone = ?", userInput.EmailOrPhone, userInput.EmailOrPhone).
			Update("is_active", true)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"status":     "error",
				"statusCode": 500,
				"message":    "Failed to activate account",
				"errors":     result.Error.Error(),
			})
			return
		}
		// Make sure user.UUID is properly set before this line
		var user models.User

		// Preload "Otp" relation and find the user by UUID
		if err := config.DB.Preload("Otp").Where("email = ? OR phone = ?", userInput.EmailOrPhone, userInput.EmailOrPhone).First(&user).Error; err != nil {
			// Handle error if user is not found or query fails
			c.JSON(500, gin.H{
				"status":     "error",
				"statusCode": 500,
				"message":    "Failed to reload user with OTP: " + err.Error(),
			})
			return
		}

		// Set the Location header to indicate the newly created user URL
		c.Header("Location", "/db/v1/auth/register/"+user.Email)

		// Respond with the user data and success message
		c.JSON(201, gin.H{
			"status":     "success",
			"statusCode": 201,
			"data":       user,
			"message":    "User activation successfully",
		})

	} else {
		// Jika OTP tidak valid
		c.JSON(400, gin.H{
			"status":     "error",
			"statusCode": 400,
			"message":    "Invalid OTP",
		})
	}

}

func ResendActivationCodeController(c *gin.Context) {

	var userInput models.ResendActivationCode

	// Ambil Bearer Token dari header Authorization
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(400, gin.H{
			"status":     "error",
			"statusCode": 400,
			"message":    "Authorization token is required",
		})
		return
	}

	// Verifikasi token (misalnya dengan middleware atau helper function)
	// Jika perlu verifikasi token, bisa dipanggil di sini
	// if !helpers.ValidateToken(token) {
	//   c.JSON(401, gin.H{
	//     "status":    "error",
	//     "message":   "Unauthorized",
	//     "statusCode": 401,
	//   })
	//   return
	// }

	// Bind input multipart-form data ke struct
	if err := c.ShouldBind(&userInput); err != nil {
		c.JSON(422, gin.H{
			"status":     "error",
			"statusCode": 422,
			"message":    "Invalid input data",
			"errors":     err.Error(),
		})
		return
	}

	// Panggil service untuk mengirimkan OTP
	success, err := services.ResendActivationCodeService(&userInput)
	if err != nil {
		// Jika terjadi error saat memproses permintaan
		c.JSON(500, gin.H{
			"status":     "error",
			"statusCode": 500,
			"message":    "Failed to resend activation code",
			"errors":     err.Error(),
		})
		return
	}

	// Mengembalikan response sukses jika OTP berhasil dikirim
	c.JSON(201, gin.H{
		"status":     "success",
		"statusCode": 201,
		"message":    "Activation code sent successfully",
		"success":    success,
	})
}

func EmailConfirmController(c *gin.Context) {
	c.String(200, "Hello PRENS ResendActivationCode")
}
func ResendEmailConfirmController(c *gin.Context) {
	c.String(200, "Hello PRENS ResendEmailConfirmController")
}
