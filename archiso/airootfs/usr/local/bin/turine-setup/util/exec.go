package util

import (
	"os"
	"os/exec"
	"strings"
)

func Run(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin // important for sudo password
	return c.Run()
}

func WhichPackage(pkg string)  error{
	return Run("which", pkg)
}

func MakepkgInstall() error {
	return Run("makepkg", "-si", "--noconfirm")
}

func EnableMultilib() error {
	data, err := os.ReadFile("/etc/pacman.conf")
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	inMultilib := false

	for i, line := range lines {
		switch {
		case strings.HasPrefix(line, "#[multilib]"):
			lines[i] = "[multilib]"
			inMultilib = true

		case inMultilib && strings.HasPrefix(line, "#Include"):
			lines[i] = strings.TrimPrefix(line, "#")
			inMultilib = false
		}
	}

	if err := os.WriteFile(
		"/etc/pacman.conf",
		[]byte(strings.Join(lines, "\n")),
		0644,
	); err != nil {
		return err
	}

	return Run("sudo", "pacman", "-Sy")
}

func IsJackInstalled() (bool, error) {
	cmd := exec.Command("pacman", "-Qq", "jack", "jack2")
	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Run()
	if err == nil {
		// at least one of them is installed
		return true, nil
	}

	// pacman returns exit code 1 if none are installed
	if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
		return false, nil
	}

	return false, err
}
