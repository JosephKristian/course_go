package models

type RegisterInput struct {
	Name                string `json:"name" form:"name" query:"name"`
	Email               string `json:"email" form:"email" query:"email"`
	Phone               string `json:"phone" form:"phone" query:"phone"`
	Password            string `json:"password" form:"password" query:"password"`
	DeviceID            string `json:"device_id" form:"device_id" query:"device_id"`
	DeviceName          string `json:"device_name" form:"device_name" query:"device_name"`
	Latitude            string `json:"latitude" form:"latitude" query:"latitude"`
	Longitude           string `json:"longitude" form:"longitude" query:"longitude"`
	ReferralCode        string `json:"referral_code" form:"referral_code" query:"referral_code"`
	Website             string `json:"website" form:"website" query:"website"`
	VerificationChannel string `json:"verification_channel" form:"verification_channel" query:"verification_channel"`
}
