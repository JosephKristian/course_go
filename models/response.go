package models

type ResendActivationCodeResponse struct {
	Success bool   `json:"success"`
	Otp     string `json:"otp"`
}
