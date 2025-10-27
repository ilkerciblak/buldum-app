package authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	ID    uuid.UUID `json:"user_id,omitempty"`
	Email string    `json:"email,omitempty"`
	Role  string    `json:"role,omitempty"`
	jwt.RegisteredClaims
}

func NewUserClaims(userid uuid.UUID, role string, email string, duration time.Duration) (*UserClaims, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &UserClaims{
		Email: email,
		Role:  role,
		ID:    userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenId.String(),
			Subject:   email,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, nil
}
