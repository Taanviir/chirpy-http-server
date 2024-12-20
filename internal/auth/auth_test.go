package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCheckPasswordHash(t *testing.T) {
	// First, we need to create some hashed passwords for testing
	password1 := "correctPassword123!"
	password2 := "anotherPassword456!"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: password1,
			hash:     hash1,
			wantErr:  false,
		},
		{
			name:     "Incorrect password",
			password: "wrongPassword",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Password doesn't match different hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Invalid hash",
			password: password1,
			hash:     "invalidhash",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMakeJWT(t *testing.T) {
	tokenSecret := "my_secret_key"
	userID := uuid.New()
	expiresIn := time.Hour

	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	assert.NoError(t, err, "Error creating JWT token")
	assert.NotEmpty(t, token, "Token should not be empty")
}

func TestValidateJWT(t *testing.T) {
	tokenSecret := "my_secret_key"
	userID := uuid.New()
	expiresIn := time.Hour

	// Generate a token to validate
	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	assert.NoError(t, err, "Error creating JWT token")

	// Validate the token
	validatedUserID, err := ValidateJWT(token, tokenSecret)
	assert.NoError(t, err, "Error validating JWT token")
	assert.Equal(t, userID, validatedUserID, "Validated user ID should match the original")
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	tokenSecret := "my_secret_key"
	userID := uuid.New()
	expiresIn := -time.Hour // Token is already expired

	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	assert.NoError(t, err, "Error creating JWT token")

	_, err = ValidateJWT(token, tokenSecret)
	assert.Error(t, err, "Expected error for expired token")
}

func TestValidateJWT_InvalidSecret(t *testing.T) {
	tokenSecret := "my_secret_key"
	invalidSecret := "wrong_secret_key"
	userID := uuid.New()
	expiresIn := time.Hour

	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	assert.NoError(t, err, "Error creating JWT token")

	_, err = ValidateJWT(token, invalidSecret)
	assert.Error(t, err, "Expected error for invalid secret")
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	tokenSecret := "my_secret_key"
	invalidToken := "this.is.not.a.valid.jwt"

	_, err := ValidateJWT(invalidToken, tokenSecret)
	assert.Error(t, err, "Expected error for invalid token")
}

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		wantToken string
		wantErr   string
	}{
		{
			name:      "No Authorization Header",
			headers:   http.Header{},
			wantToken: "",
			wantErr:   "authorization header is missing",
		},
		{
			name: "Authorization Header Without Bearer Prefix",
			headers: http.Header{
				"Authorization": []string{"Token abc123"},
			},
			wantToken: "",
			wantErr:   "missing Bearer token",
		},
		{
			name: "Valid Bearer Token",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			wantToken: "abc123",
			wantErr:   "",
		},
		{
			name: "Bearer Token With Leading Spaces",
			headers: http.Header{
				"Authorization": []string{"  Bearer   abc123"},
			},
			wantToken: "",
			wantErr:   "missing Bearer token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GetBearerToken(tt.headers)
			if token != tt.wantToken {
				t.Errorf("expected token: %q, got: %q", tt.wantToken, token)
			}
			if (err != nil && err.Error() != tt.wantErr) || (err == nil && tt.wantErr != "") {
				t.Errorf("expected error: %q, got: %v", tt.wantErr, err)
			}
		})
	}
}
