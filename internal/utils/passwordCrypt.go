package utils

import (
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

func HashFromPassword(passwordString string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(
			"failed to hash password",
			slog.String("error", err.Error()),
			slog.String("context", "HashFromPassword"),
		)
		return nil
	}
	return hash
}
