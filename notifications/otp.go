package notification

import (
	"errors"
	"fmt"
	"net/smtp"
)

func SendEmail(to, name, processName, otp, lang string) error {
	// Subject email
	subject := fmt.Sprintf("OTP untuk %s: %s", processName, otp)

	// Body email
	body := fmt.Sprintf("OTP Anda: %s\nBerhenti sebentar: OTP akan berlaku selama beberapa menit.", otp)

	// Membuat header email
	headers := fmt.Sprintf("From: %s\r\n", "no-reply@example.com")
	headers += fmt.Sprintf("To: %s\r\n", to)
	headers += fmt.Sprintf("Subject: %s\r\n", subject)

	// Gabungkan header dan body menjadi satu pesan
	message := []byte(headers + "\r\n" + body)

	// Kirim email menggunakan smtp.SendMail
	return smtp.SendMail("smtp.example.com:587", nil, "no-reply@example.com", []string{to}, message)
}

// Fungsi untuk mengirim SMS
func SendSms(receiver string, message string) error {
	// Misalnya, menggunakan layanan SMS API di sini.
	// Simulasikan pengiriman SMS
	if receiver == "" || message == "" {
		return errors.New("receiver or message is empty")
	}

	// Logika pengiriman SMS
	fmt.Printf("Sending SMS to %s: %s\n", receiver, message)

	// Return nil error jika pengiriman berhasil
	return nil
}

// Fungsi untuk mengirim WhatsApp
func SendWhatsapp(receiver string, processName string, otp string, expiredInMinutes int, lang string) error {
	// WhatsApp API di sini.
	// Simulasikan pengiriman WhatsApp
	if receiver == "" || otp == "" {
		return errors.New("receiver or OTP is empty")
	}

	// Buat pesan WhatsApp
	message := fmt.Sprintf("OTP untuk %s: %s. OTP berlaku selama %d menit.", processName, otp, expiredInMinutes)

	// Logika pengiriman WhatsApp
	fmt.Printf("Sending WhatsApp message to %s: %s\n", receiver, message)

	// Return nil error jika pengiriman berhasil
	return nil
}
