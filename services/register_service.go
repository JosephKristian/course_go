package services

// import (
// 	"errors"
// 	"fmt"
// 	"strings"
// 	"time"

// 	"github.com/google/uuid"
// 	"gitlab.com/ipaymupreviews/golang-gin-poc/helpers"
// 	"gitlab.com/ipaymupreviews/golang-gin-poc/models"
// )

// func registerService(data *models.RegisterInput) (*models.User, error) {
// 	// Validasi input
// 	if data.Name == "" {
// 		return nil, errors.New("name is required")
// 	}

// 	if data.Email != "" && !helpers.IsValidEmail(data.Email) {
// 		return nil, errors.New("invalid email format")
// 	}

// 	isEmailExist, err := r.userRepo.CheckEmailExist(data.Email)
// 	if err != nil {
// 		return nil, errors.New("failed to check email existence: " + err.Error())
// 	}
// 	if isEmailExist {
// 		return nil, errors.New("email already registered")
// 	}

// 	if data.Phone == "" {
// 		return nil, errors.New("phone is required")
// 	}
// 	if len(data.Phone) < 8 || len(data.Phone) > 15 {
// 		return nil, errors.New("phone number must be between 8 and 15 digits")
// 	}

// 	if data.Password == "" {
// 		return nil, errors.New("password is required")
// 	} else if len(data.Password) < 8 || !helpers.ContainsRequiredChars(data.Password) {
// 		return nil, errors.New("password must be at least 8 characters, include uppercase, lowercase, number, and special character")
// 	}

// 	if data.VerificationChannel == "" {
// 		return nil, errors.New("verification_channel is required")
// 	}

// 	// Hashing password
// 	hashedPassword, err := helpers.HashPassword(data.Password)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to hash password: %v", err)
// 	}

// 	// UUID dan OTP
// 	userUUID := uuid.NewString()

// 	// Simpan data ke database
// 	user := &models.User{
// 		Name:             data.Name,
// 		Email:            strings.ToLower(strings.TrimSpace(data.Email)),
// 		Phone:            helpers.FormatPhoneNumber(data.Phone),
// 		Password:         hashedPassword,
// 		UUID:             userUUID,
// 		DeviceID:         data.DeviceID,
// 		DeviceName:       data.DeviceName,
// 		ConfirmationFlow: data.VerificationChannel,
// 		CreatedAt:        time.Now(),
// 		UpdatedAt:        time.Now(),
// 	}

// 	if err := r.userRepo.Save(user); err != nil {
// 		return nil, fmt.Errorf("failed to save user: %v", err)
// 	}

// 	// Kirim OTP melalui OtpService
// 	processName := "User Registration"
// 	token := userUUID
// 	receiverName := data.Name
// 	receiver := data.Email
// 	if data.VerificationChannel == "whatsapp" {
// 		receiver = data.Phone
// 	}

// 	_, err = r.otpService.SendOtp(data.VerificationChannel, processName, token, receiverName, receiver, 6, 300, "en")
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to send verification: %v", err)
// 	}

// 	return user, nil
// }
