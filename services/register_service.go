package services

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"gitlab.com/ipaymupreviews/golang-gin-poc/config"
	"gitlab.com/ipaymupreviews/golang-gin-poc/helpers"
	"gitlab.com/ipaymupreviews/golang-gin-poc/models"
	"gorm.io/gorm"
)

func ResendActivationCodeService(data *models.ResendActivationCode) (*models.ResendActivationCodeResponse, error) {
	const processName = "account_activation"

	// Ambil data user berdasarkan email atau nomor telepon
	var user models.User
	err := config.DB.
		Where("email = ? OR phone = ?", data.EmailOrPhone, data.EmailOrPhone).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to retrieve user: %v", err)
	}

	// Generate OTP
	otp := helpers.GenerateOtp(6)
	expiredAt := time.Now().Add(300 * time.Second) // 5 menit

	// Update OTP dan expired_at di tabel otps
	err = config.DB.
		Table("otps").
		Where("user_id = ?", user.UUID).
		Updates(map[string]interface{}{
			"otp":        otp,
			"expired_at": expiredAt,
		}).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update OTP: %v", err)
	}

	// Kirim OTP
	_, err = helpers.SendOtp(
		otp,
		data.VerificationChannel,
		processName,
		user.UUID,
		user.Name, // Nama user
		data.EmailOrPhone,
		6,    // Panjang OTP
		300,  // Waktu kedaluwarsa OTP (dalam detik)
		"en", // Bahasa OTP
	)
	if err != nil {
		return nil, fmt.Errorf("failed to send OTP: %v", err)
	}

	response := &models.ResendActivationCodeResponse{
		Success: true,
		Otp:     otp, // Masukkan OTP ke dalam response
	}
	return response, nil
}

func RegisterService(data *models.RegisterInput) (*models.User, error) {

	var existingUser models.User

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
	if err := config.DB.Where("email = ?", data.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to check email existence: %v", err)
	}

	formattedPhone := helpers.FormatPhoneNumber(data.Phone) // Pastikan format nomor telepon sudah benar
	if err := config.DB.Where("phone = ?", formattedPhone).First(&existingUser).Error; err == nil {
		return nil, errors.New("phone number already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to check phone number existence: %v", err)
	}

	// Hashing password
	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Buat data user
	user := &models.User{
		Name:             data.Name,
		Email:            strings.ToLower(strings.TrimSpace(data.Email)),
		Phone:            helpers.FormatPhoneNumber(data.Phone),
		Password:         hashedPassword,
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

	// Generate OTP
	otp := helpers.GenerateOtp(6)
	expiredAt := time.Now().Add(time.Duration(300) * time.Second)

	if err := config.DB.Where("uuid = ?", user.UUID).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user with token %s not found: %w", user.UUID, err)
	}

	otpData := models.Otp{
		UserID:      user.UUID,
		Otp:         otp,
		Destination: receiver,
		Flow:        processName,
		Channel:     data.VerificationChannel,
		ExpiredAt:   expiredAt,
	}

	// Simpan data OTP ke database
	if err := config.DB.Create(&otpData).Error; err != nil {
		return nil, fmt.Errorf("failed to save OTP to database: %w", err)
	}

	_, err = helpers.SendOtp(otp,
		data.VerificationChannel,
		processName,
		user.UUID,
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

	// Memuat ulang objek user dengan OTP terbaru
	if err := config.DB.Preload("Otp").Where("uuid = ?", user.UUID).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to reload user with OTP: %v", err)
	}

	return user, nil
}
func VerifyOtp(emailOrPhone string, verificationCode int) (bool, error) {
	var otpData models.Otp

	// Query untuk mencari OTP berdasarkan email atau phone
	err := config.DB.
		Table("otps").
		Joins("JOIN users ON users.uuid = otps.user_id").
		Where("users.email = ? OR users.phone = ?", emailOrPhone, emailOrPhone).
		First(&otpData).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("OTP not found")
		}
		return false, fmt.Errorf("failed to retrieve OTP: %v", err)
	}

	// Cek apakah OTP sudah kedaluwarsa
	if time.Now().After(otpData.ExpiredAt) {
		return false, errors.New("OTP expired")
	}

	// Cek apakah kode OTP sesuai dengan yang diberikan
	// Pastikan OTP yang ada dalam database sesuai dengan tipe data yang Anda harapkan
	// Jika otpData.Otp bertipe string, Anda harus membandingkannya dengan string yang diformat
	if fmt.Sprintf("%06d", verificationCode) != otpData.Otp {
		return false, errors.New("invalid verification code")
	}

	// Jika verifikasi berhasil, update status OTP menjadi terverifikasi
	err = config.DB.Model(&otpData).Update("verified_at", time.Now()).Error
	if err != nil {
		return false, fmt.Errorf("failed to update OTP verification status: %v", err)
	}

	log.Printf("[INFO] OTP verified successfully for: %s", emailOrPhone)
	return true, nil
}
