package auth

import (
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

type Claim struct {
	jwt.StandardClaims
}

type SignConfig struct {
	Key        string
	ValidUntil time.Duration
}

func SignJWT(userId int64, conf SignConfig) (string, error) {
	claim := Claim{
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.FormatInt(userId, 10),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(conf.ValidUntil).Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := t.SignedString([]byte(conf.Key))
	if err != nil {
		return "", err
	}
	return ss, nil
}
