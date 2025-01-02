package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID             string    `json:"uuid" form:"uuid" gorm:"primary_key"`
	Name             string    `json:"name" form:"name" validate:"required"`
	Email            string    `json:"email" form:"email" validate:"omitempty,email"`
	Phone            string    `json:"phone" form:"phone" validate:"required,numeric,len=10"`
	Password         string    `json:"password" form:"password" validate:"omitempty,min=8"`
	ReferralCode     string    `json:"referral_code" form:"referral_code" validate:"omitempty"`
	Website          string    `json:"website" form:"website" validate:"omitempty,url"`
	ConfirmationFlow string    `json:"confirmation_flow" form:"confirmation_flow" validate:"omitempty,oneof=phone email"`
	DeviceID         string    `json:"device_id" form:"device_id"`
	DeviceName       string    `json:"device_name" form:"device_name"`
	IP               string    `json:"ip,omitempty" form:"ip,omitempty"`
	VerificationCode string    `json:"verification_code" form:"verification_code"` // Added VerificationCode field
	CreatedAt        time.Time `json:"created_at" form:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" form:"updated_at"`
}
