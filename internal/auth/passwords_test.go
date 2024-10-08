// test_passwords.go
package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword returned an error: %v", err)
	}

	if len(hash) == 0 {
		t.Errorf("HashPassword returned an empty hash")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecretpassword"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword returned an error: %v", err)
	}

	err = CheckPasswordHash(password, hash)
	if err != nil {
		t.Errorf("CheckPasswordHash returned an error: %v", err)
	}
}

func TestCheckPasswordHashInvalidPassword(t *testing.T) {
	password := "mysecretpassword"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword returned an error: %v", err)
	}

	err = CheckPasswordHash("wrongpassword", hash)
	if err == nil {
		t.Errorf("CheckPasswordHash did not return an error for an invalid password")
	}
}

func TestCheckPasswordHashInvalidHash(t *testing.T) {
	password := "mysecretpassword"
	err := CheckPasswordHash(password, "invalidhash")
	if err == nil {
		t.Errorf("CheckPasswordHash did not return an error for an invalid hash")
	}
}
