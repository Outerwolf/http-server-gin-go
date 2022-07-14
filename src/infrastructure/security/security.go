package security

import "github.com/golang-jwt/jwt"

var secret = []byte("secret")

type ITokenServices interface {
	SignToken(cred *Credentials) (string, error)
	VerifyToken(tokenString string) (bool, error)
}

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

func SignToken(cred *Credentials) (string, error) {
	claims := &Claims{
		UserName: cred.UserName,
		Email:    cred.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
