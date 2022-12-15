package domains

import "time"

type User struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	EmailVerification time.Time `json:"email_verification"`
	Password          string    `json:"password"`
	RememberToken     string    `json:"remember_token"`
	Slug              string    `json:"slug"`
	Gender            string    `json:"gender"`
	Relationship      string    `json:"relationship"`
	Job               string    `json:"job"`
	Company           string    `json:"company"`
	Phone             string    `json:"phone"`
	Status            string    `json:"status"`
	Role              string    `json:"role"`
	FirebaseUUID      string    `json:"firebase_uuid"`
	ResetPasswordCode string    `json:"reset_password_code"`
	IsVerified        int       `json:"is_verified"`
	VerificationCode  string    `json:"verification_code"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
