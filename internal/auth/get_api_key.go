package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (key string, err error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	prefix := "ApiKey "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("missing ApiKey token")
	}

	return strings.TrimPrefix(authHeader, prefix), nil
}
