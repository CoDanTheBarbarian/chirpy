package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "mysecret"
	expiresIn := 1 * time.Hour

	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Errorf("MakeJWT returned an error: %v", err)
	}

	validatedUserID, err := ValidateJWT(token, tokenSecret)
	if err != nil {
		t.Errorf("ValidateJWT returned an error: %v", err)
	}

	if validatedUserID != userID {
		t.Errorf("Expected user ID %v, got %v", userID, validatedUserID)
	}
}

func TestValidateExpiredJWT(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "mysecret"
	expiresIn := -1 * time.Hour // token is already expired

	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Errorf("MakeJWT returned an error: %v", err)
	}

	_, err = ValidateJWT(token, tokenSecret)
	if err == nil {
		t.Errorf("ValidateJWT did not return an error for an expired token")
	}
}

func TestValidateJWTWithWrongSecret(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "mysecret"
	wrongSecret := "wrongsecret"
	expiresIn := 1 * time.Hour

	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Errorf("MakeJWT returned an error: %v", err)
	}

	_, err = ValidateJWT(token, wrongSecret)
	if err == nil {
		t.Errorf("ValidateJWT did not return an error for a token signed with the wrong secret")
	}
}
