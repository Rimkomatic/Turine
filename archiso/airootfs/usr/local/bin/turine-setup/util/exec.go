package util

import (
	"os"
	"os/exec"
)

func Run(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin // important for sudo password
	return c.Run()
}

func PacmanInstall(pkgs ...string) error {
	args := append([]string{"pacman", "-S", "--noconfirm"}, pkgs...)
	return Run("sudo", args...)
}
