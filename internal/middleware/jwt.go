package middleware

import (
	"errors"
	cfg "go-mitra-akosta/pkg/utils/config"
	"go-mitra-akosta/pkg/utils/logger"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(id int, email, role string) string
	ValidateToken(encode string) (*jwt.Token, error)
}

type claims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type JwtServices struct {
	SecretKey string
	Issuer    string
	Log       *logger.Logging
}

func JWTAuthService(cfg *cfg.Config, log *logger.Logging) JWTService {
	return &JwtServices{
		SecretKey: cfg.Jwt.Secret,
		Issuer:    "go-boilerplate",
		Log:       log,
	}
}

func (s *JwtServices) GenerateToken(id int, email, role string) string {

	claims := &claims{
		id,
		email,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Unix(time.Now().Add(time.Hour*24).Unix(), 0)},
			// ExpiresAt: &jwt.NumericDate{Time: time.Unix(time.Now().Add(time.Hour * 24), 0)},
			Issuer:   s.Issuer,
			IssuedAt: &jwt.NumericDate{Time: time.Unix(time.Now().Unix(), 0)},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encrypt, error := token.SignedString(token)

	if error != nil {
		s.Log.FatalLog("Error Encrypt Token: ", error)
	}
	return encrypt

}

func (s *JwtServices) ValidateToken(encodeToken string) (*jwt.Token, error) {
	return jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("invalid token")
		}
		return []byte(s.SecretKey), nil
	})
}
