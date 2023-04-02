package token

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/laminne/qip/pkg/utils/id"
)

type JWTTokenGenerator struct {
	key string
}

func NewJWTTokenGenerator(key string) *JWTTokenGenerator {
	return &JWTTokenGenerator{key: key}
}

func (g *JWTTokenGenerator) NewToken(uid id.SnowFlakeID) (string, error) {
	c := jwt.RegisteredClaims{
		Subject:   string(uid),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	res, err := t.SignedString([]byte(g.key))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return res, nil
}
