package auth

import (
	"fmt"
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

func ValidateJWT(tokenString string, conf SignConfig) (*Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.Key), nil
	})
	if err != nil {
		return nil, err
	} else if claim, ok := token.Claims.(*Claim); ok && token.Valid {
		return claim, nil
	}
	return nil, fmt.Errorf("invalid token")

}
