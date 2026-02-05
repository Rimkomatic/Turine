package util


func PacmanInstall(pkgs ...string) error {
	args := append([]string{"pacman", "-S", "--noconfirm"}, pkgs...)
	return Run("sudo", args...)
}

func PacmanInstallU(pkgs ...string) error {
	args := append([]string{"pacman", "-U", "--noconfirm"}, pkgs...)
	return Run("sudo", args...)
}
