package util

import "strings"

func StowDotfiles(pkgs string) error {
	args := strings.Fields(pkgs)
	return Run("stow", args...)
}
