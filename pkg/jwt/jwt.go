package jwt

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	Secret string
}

func NewJWT(secter string) *JWT {
	return &JWT{
		Secret: secter,
	}
}

func (j *JWT) Create(email string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email": email,
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

