package helpers

import (
	"math/rand"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// ContainsRequiredChars mengecek apakah password memiliki kombinasi karakter tertentu.
func ContainsRequiredChars(password string) bool {
	// Memeriksa apakah password mengandung setidaknya satu huruf kecil
	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false

	for _, c := range password {
		switch {
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsDigit(c):
			hasDigit = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	// Semua kondisi harus terpenuhi
	return hasLower && hasUpper && hasDigit && hasSpecial
}

// isValidEmail memvalidasi format email.
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// generateOtp menghasilkan kode OTP dengan panjang tertentu.
func GenerateOtp(length int) string {
	otp := make([]byte, length)
	for i := range otp {
		otp[i] = byte(rand.Intn(10) + '0')
	}
	return string(otp)
}

// generateRandomPassword menghasilkan password acak dengan panjang tertentu.
func GenerateRandomPassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// hashPassword melakukan hashing pada password.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// formatPhoneNumber membersihkan format nomor telepon.
func FormatPhoneNumber(phone string) string {
	return strings.ReplaceAll(phone, " ", "")
}
