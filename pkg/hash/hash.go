package hash

import (
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to generate hash from password", "error", err)
		return ""
	}
	return string(hash)
}

func CompareHashPassword(hashedPassword, requestPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestPassword)); err != nil {
		return err
	}
	return nil
}
