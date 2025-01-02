package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.com/ipaymupreviews/golang-gin-poc/config"
	"gitlab.com/ipaymupreviews/golang-gin-poc/helpers"
	"gitlab.com/ipaymupreviews/golang-gin-poc/models"
	"gorm.io/gorm"
)

func registerService(data *models.RegisterInput) (*models.User, error) {
	// Validasi input
	if data.Name == "" {
		return nil, errors.New("name is required")
	}

	if data.Email != "" && !helpers.IsValidEmail(data.Email) {
		return nil, errors.New("invalid email format")
	}

	if data.Phone == "" {
		return nil, errors.New("phone is required")
	}
	if len(data.Phone) < 8 || len(data.Phone) > 15 {
		return nil, errors.New("phone number must be between 8 and 15 digits")
	}

	if data.Password == "" {
		return nil, errors.New("password is required")
	} else if len(data.Password) < 8 || !helpers.ContainsRequiredChars(data.Password) {
		return nil, errors.New("password must be at least 8 characters, include uppercase, lowercase, number, and special character")
	}

	if data.VerificationChannel == "" {
		return nil, errors.New("verification_channel is required")
	}

	// Validasi email unik
	var existingUser models.User
	if err := config.DB.Where("email = ?", data.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to check email existence: %v", err)
	}

	// Hashing password
	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// UUID untuk user
	userUUID := uuid.NewString()

	// Buat data user
	user := &models.User{
		Name:             data.Name,
		Email:            strings.ToLower(strings.TrimSpace(data.Email)),
		Phone:            helpers.FormatPhoneNumber(data.Phone),
		Password:         hashedPassword,
		UUID:             userUUID,
		DeviceID:         data.DeviceID,
		DeviceName:       data.DeviceName,
		ConfirmationFlow: data.VerificationChannel,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	// Simpan data user ke database
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to save user: %v", err)
	}

	// Kirim OTP melalui OtpService
	processName := "User Registration"
	receiver := data.Email
	if data.VerificationChannel == "whatsapp" {
		receiver = data.Phone
	}

	_, err = helpers.SendOtp(data.VerificationChannel,
		processName,
		userUUID,
		data.Name,
		receiver,
		6,    // Panjang OTP
		300,  // Waktu kedaluwarsa OTP (dalam detik)
		"en", // Bahasa OTP
	)

	if err != nil {
		// Hapus user jika pengiriman OTP gagal
		config.DB.Delete(&user)
		return nil, fmt.Errorf("failed to send verification: %v", err)
	}

	return user, nil
}
