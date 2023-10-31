package middleware

import (
	"errors"
	"job-portal-api/internal/auth"
)

type Mid struct {
	auth *auth.Auth
}

func NewMid(a *auth.Auth) (Mid, error) {
	if a == nil {
		return Mid{}, errors.New("auth can't be nil")
	}

	return Mid{auth: a}, nil
}
