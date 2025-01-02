package models

import (
	"time"

	"gorm.io/gorm"
)

// Otp represents the structure of an OTP entity.
type Otp struct {
	gorm.Model
	ID          uint       `json:"id" form:"id" gorm:"primaryKey"`                     // Auto-increment primary key (if using GORM)
	Uuid        string     `json:"uuid" form:"uuid" gorm:"uniqueIndex"`                // Unique identifier
	Otp         string     `json:"otp" form:"otp"`                                     // OTP code
	Token       string     `json:"token" form:"token"`                                 // Token/session identifier
	Destination string     `json:"destination" form:"destination"`                     // Target destination (e.g., email, phone)
	Flow        string     `json:"flow" form:"flow"`                                   // Process name (e.g., "registration")
	Channel     string     `json:"channel" form:"channel"`                             // Delivery method (e.g., email, SMS)
	ExpiredAt   time.Time  `json:"expired_at" form:"expired_at"`                       // Expiration time
	CreatedAt   time.Time  `json:"created_at" form:"-"`                                // Auto-set by GORM, not exposed in form
	UpdatedAt   time.Time  `json:"updated_at" form:"-"`                                // Auto-set by GORM, not exposed in form
	VerifiedAt  *time.Time `json:"verified_at" form:"verified_at" gorm:"default:null"` // Nullable, set upon verification
}
