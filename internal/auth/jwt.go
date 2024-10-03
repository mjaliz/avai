package auth

import "github.com/golang-jwt/jwt"

type Claim struct {
	jwt.StandardClaims
}

func SignJWT(claim *Claim, key string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := t.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return ss, nil
}
