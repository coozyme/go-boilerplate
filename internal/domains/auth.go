package domains

import (
	"time"
)

type Roles struct {
	Id   int
	Name string
}

type Auth struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	EmailVerification time.Time `json:"email_verification"`
	Password          string    `json:"password"`
	RememberToken     string    `json:"remember_token"`
	Role              string    `json:"role"`
	FirebaseUUID      string    `json:"firebase_uuid"`
	ResetPasswordCode string    `json:"reset_password_code"`
	IsVerified        int       `json:"is_verified"`
	VerificationCode  string    `json:"verification_code"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type AuthLogin struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	IsVerified int    `json:"is_verified"`
}

type AuthRegister struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   string `json:"status"`
	Slug     string `json:"slug"`
}

type AuthRepository interface {
	FindAccount(email string) bool
}

type AuthUsecase interface {
	CheckAccount(email string) bool
	Login(email, password string) (*AuthLogin, error)
	Register(user *AuthRegister) (*Auth, error)
}
