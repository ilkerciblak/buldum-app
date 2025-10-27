package authentication

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTHelper struct {
	secretKey     string
	signingMethod jwt.SigningMethod
}

func NewJWTMaker(secretKey string) *JWTHelper {
	return &JWTHelper{
		secretKey: secretKey,
	}
}

func (j *JWTHelper) CreateToken(userid uuid.UUID, role string, email string, duration time.Duration) (string, *UserClaims, error) {
	claims, err := NewUserClaims(userid, role, email, duration)
	if err != nil {
		return "", nil, err
	}

	token := jwt.NewWithClaims(j.signingMethod, claims)
	tokenStr, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", nil, err
	}

	return tokenStr, claims, nil
}

func (j *JWTHelper) ValidateToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(t *jwt.Token) (any, error) {

		if t.Method.Alg() != j.signingMethod.Alg() {
			return nil, fmt.Errorf("Invalid token signing method")
		}

		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
