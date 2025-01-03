package helpers

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	notification "gitlab.com/ipaymupreviews/golang-gin-poc/notifications"
)

func GenerateOtp(length int) string {
	rand.Seed(time.Now().UnixNano())
	otp := make([]byte, length)
	for i := range otp {
		otp[i] = byte('0' + rand.Intn(10))
	}
	return string(otp)
}

// SendOtp sends an OTP to the specified channel (email, sms, or whatsapp).
func SendOtp(otp, via, processName, token, receiverName, receiver string, otpLength, expiredInSeconds int, lang string) (string, error) {
	var sendErr error
	var otpSentMessage string

	switch strings.ToLower(via) {
	case "email":
		sendErr = notification.SendEmail(receiver, receiverName, processName, otp, lang)
		otpSentMessage = "OTP sent successfully via email."
	case "sms":
		message := fmt.Sprintf("Your OTP for %s is %s. It is valid for %d minutes.", processName, otp, expiredInSeconds/60)
		sendErr = notification.SendSms(receiver, message)
		otpSentMessage = "OTP sent successfully via SMS."
	case "whatsapp":
		sendErr = notification.SendWhatsapp(receiver, processName, otp, expiredInSeconds/60, lang)
		otpSentMessage = "OTP sent successfully via WhatsApp."
	default:
		return "", fmt.Errorf("unsupported OTP channel: %s", via)
	}

	// If there's an error while sending the OTP
	if sendErr != nil {
		log.Printf("[ERROR] Failed to send OTP via %s to %s: %v", via, receiver, sendErr)
		return "", fmt.Errorf("failed to send OTP via %s: %w", via, sendErr)
	}

	log.Printf("[INFO] %s to %s.", otpSentMessage, receiver)
	return otpSentMessage, nil
}
