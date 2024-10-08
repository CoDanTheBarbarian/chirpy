package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	byte_array := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(byte_array, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) error {
	byte_array := []byte(password)
	byte_hash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byte_hash, byte_array)
	if err != nil {
		return err
	}
	return nil
}
