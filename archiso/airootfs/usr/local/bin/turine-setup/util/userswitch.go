package util

import (
	"errors"
	"os"
)

func RunAsUser(cmd string, args ...string) error {
	if sudoUser := os.Getenv("SUDO_USER"); sudoUser != "" {
		fullArgs := append([]string{"-iu", sudoUser, cmd}, args...)
		return Run("sudo", fullArgs...)
	}

	if os.Geteuid() != 0 {
		return Run(cmd, args...)
	}

	return errors.New("cannot determine non-root user context")
}

