package middleware

import (
	"context"
	"net/http"
	"strings"

	auth "github.com/ilkerciblak/buldum-app/shared/authentication"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

type AuthenticationMiddleware struct {
	logging.ILogger
	auth.JWTHelper
}

func (m *AuthenticationMiddleware) Act(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr, err := m.getAuthorizationHeader(r)
		if err != nil {
			corepresentation.RespondWithProblemDetails(w, corepresentation.NewErrorResult(err).Error)
			return
		}
		claims, err := m.verifyToken(r, tokenStr)
		if err != nil {
			corepresentation.RespondWithProblemDetails(w, corepresentation.NewErrorResult(err).Error)
			return
		}
		type authKey struct{}
		ctx := context.WithValue(r.Context(), authKey{}, claims)
		handlerFunc.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (a AuthenticationMiddleware) getAuthorizationHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		a.ILogger.Log(logging.DEBUG, r.Context(), "", "error", "authorization header is missing")
		return "", coredomain.UserNotAuthenticated.WithMessage("Authorization header is missing")
	}
	fields := strings.Fields(authHeader)
	if len(fields) != 2 || fields[0] != "Bearer" {
		return "", coredomain.UserNotAuthenticated.WithMessage("Invalid authorization header")
	}

	return fields[1], nil
}

func (a AuthenticationMiddleware) verifyToken(r *http.Request, tokenStr string) (*auth.UserClaims, error) {
	claims, err := a.JWTHelper.ValidateToken(tokenStr)
	if err != nil {
		a.ILogger.Log(logging.DEBUG, r.Context(), "", "error", err.Error())
		return nil, coredomain.UserNotAuthenticated.WithMessage("Invalid authorization token")
	}

	return claims, nil

}
