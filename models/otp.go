package models

import (
	"time"

	"gorm.io/gorm"
)

type Otp struct {
	gorm.Model
	UUID        string     `json:"uuid" form:"uuid" gorm:"default:uuid_generate_v4();"`
	Otp         string     `json:"otp" form:"otp"`
	Token       string     `json:"token" form:"token"`
	Destination string     `json:"destination" form:"destination"`
	Flow        string     `json:"flow" form:"flow"`
	Channel     string     `json:"channel" form:"channel"`
	ExpiredAt   time.Time  `json:"expired_at" form:"expired_at"`
	VerifiedAt  *time.Time `json:"verified_at" form:"verified_at" gorm:"default:null"`
	CreatedAt   time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" form:"updated_at"`

	UserID string `json:"user_id" form:"user_id" gorm:"index"`
}
