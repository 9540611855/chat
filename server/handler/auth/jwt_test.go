package auth

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestGenerateAndValidateJWT(t *testing.T) {
	userID := int64(123)
	tokenString, err := GenerateJWT(userID)
	if err != nil {
		t.Fatalf("failed to generate JWT: %v", err)
	}

	claims, err := ValidateJWT(tokenString)
	if err != nil {
		t.Fatalf("failed to validate JWT: %v", err)
	}

	if authorized, ok := claims["authorized"].(bool); !ok || !authorized {
		t.Errorf("expected authorized claim to be true, got %v", authorized)
	}

	if user, ok := claims["user"].(float64); !ok || int64(user) != userID {
		t.Errorf("expected user claim to be %d, got %v", userID, user)
	}
}

func TestCheckJwtTokenTime(t *testing.T) {
	tokenString, err := GenerateJWT(123)
	if err != nil {
		t.Fatalf("failed to generate JWT: %v", err)
	}

	// Test that a token generated less than 2 hours ago is valid
	if !CheckJwtTokenTime(tokenString, 123) {
		t.Errorf("expected token to be valid, got invalid")
	}

	// Test that a token generated more than 2 hours ago is invalid
	oldToken := jwt.New(jwt.SigningMethodHS256)
	oldClaims := oldToken.Claims.(jwt.MapClaims)
	oldClaims["authorized"] = true
	oldClaims["user"] = 123
	oldClaims["exp"] = time.Now().Add(-time.Hour * 3).Unix()
	oldTokenString, err := oldToken.SignedString([]byte("mHPqAeyMhQzR3RdVcN79"))
	if err != nil {
		t.Fatalf("failed to generate old token: %v", err)
	}
	if CheckJwtTokenTime(oldTokenString, 123) {
		t.Errorf("expected old token to be invalid, got valid")
	}
}
