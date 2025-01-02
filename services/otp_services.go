package services

import (
	"math/rand"
	"time"
)

func GenerateOtp(length int) string {
	rand.Seed(time.Now().UnixNano())
	otp := make([]byte, length)
	for i := range otp {
		otp[i] = byte('0' + rand.Intn(10))
	}
	return string(otp)
}

// Send OTP
func SendOtp(via, processName, token, receiverName, receiver string, otpLength, expiredInSeconds int, lang string) (string, error) {
	otp := GenerateOtp(otpLength)

	return otp, nil
	// expiredAt := time.Now().Add(time.Duration(expiredInSeconds) * time.Second)

	// otpUUID := uuid.New().String() // Generate UUID
	// // otpData := models.Otp{
	// // 	Uuid:        otpUUID,
	// // 	Otp:         otp,
	// // 	Token:       token,
	// // 	Destination: receiver,
	// // 	Flow:        processName,
	// // 	Channel:     via,
	// // 	ExpiredAt:   expiredAt,
	// // }

	// // err := s.otpRepository.StoreOtp(&otpData)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to store OTP: %v", err)
	// }

	// var message string
	// switch via {
	// case "email":
	// 	err = notification.SendEmail(receiver, receiverName, processName, otp, lang)
	// case "sms":
	// 	message = fmt.Sprintf("OTP for %s: %s. Valid for %d minutes.", processName, otp, expiredInSeconds/60)
	// 	err = notification.SendSms(receiver, message)
	// case "whatsapp":
	// 	err = notification.SendWhatsapp(receiver, processName, otp, expiredInSeconds/60, lang)
	// default:
	// 	return "", errors.New("unsupported channel")
	// }

	// if err != nil {
	// 	return "", fmt.Errorf("failed to send OTP: %v", err)
	// }

	// log.Printf("[INFO] OTP sent successfully via %s to %s.", via, receiver)
	// return otpData.Uuid, nil
}
