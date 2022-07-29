package models

import  (
	"time"

	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel       `bson:",inline"`
	FullName               string          `json:"fullname" bson:"fullname"`
	Phone                  string          `json:"phone" bson:"phone"`
	Username			   string			`json:"username" bson:"username"`
	Email                  string          `json:"email" bson:"email"`
	IdentificationType     string          `json:"identification_type" bson:"identification_type"`
	VerificationStatus		bool          `json:"verification_status" bson:"verification_status"`
	Password               string          `json:"password,omitempty" bson:"password"`
	TwoFactorAuth          bool            `json:"two_factor_auth" bson:"two_factor_auth"`
	ResetTokenVerified	   bool            `json:"password_reset_token_used" bson:"password_reset_token_used"`
	PasswordResetToken     string          `json:"password_reset_token" bson:"password_reset_token"`
	PasswordResetExpires   time.Time       `json:"password_reset_expires" bson:"password_reset_expires"`
	PasswordChangedAt      time.Time       `json:"password_changed_at" bson:"password_changed_at"`
	Image                  string          `json:"image" bson:"image"`
	TwoFactorAuthPin       string          `json:"two_factor_auth_pin" bson:"two_factor_auth_pin"`
	TwoFactorAuthPinExpiry time.Time       `json:"two_factor_auth_pin_expiry" bson:"two_factor_auth_pin_expiry"`
}

type SigninData struct {
	Email		string 	`json:"email"`
	Password	string	`json:"password"`
}

type SignupData struct {
	Fullname	string 		`json:"fullname"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
	Phone		string		`json:"phone"`
}