package helper

import (
	"time"

	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
)

type GoJWT struct {
}

func NewGoJWT() *GoJWT {
	return &GoJWT{}
}

func (j *GoJWT) CreateTokenJWT(user *domain.User) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID":    user.ID,
		"Roles":     user.Roles,
		"Verified":  user.Verified,
		"ExpiresAt": time.Now().Add(time.Hour * 48).Unix(),
	})

	fixToken, err := token.SignedString([]byte("220220"))
	if err != nil {
		log.Info("error create token jwt")
	}

	return fixToken
}
