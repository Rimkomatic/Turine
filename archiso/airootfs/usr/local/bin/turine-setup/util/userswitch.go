package util

import (
	"errors"
	"os"
)

func RunAsUser(cmd string, args ...string) error {
	user := os.Getenv("SUDO_USER")
	if user == "" {
		return errors.New("SUDO_USER not set")
	}

	fullArgs := append([]string{"-iu", user, cmd}, args...)
	return Run("sudo", fullArgs...)
}
