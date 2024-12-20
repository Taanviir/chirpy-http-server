package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	auth_header := headers.Get("Authorization")
	if auth_header == "" {
		return "", errors.New("authorization header is missing")
	}

	prefix := "Bearer "
	if !strings.HasPrefix(auth_header, prefix) {
		return "", errors.New("missing Bearer token")
	}

	return strings.TrimPrefix(auth_header, prefix), nil
}
