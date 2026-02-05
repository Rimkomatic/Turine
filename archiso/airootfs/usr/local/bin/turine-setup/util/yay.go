package util

import "os"


func YayInstall(pkg string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(wd)

	if err := GitClone("https://aur.archlinux.org/" + pkg + ".git"); err != nil {
		return err
	}

	if err := ChangeDirectory(pkg); err != nil {
		return err
	}

	return Run("makepkg", "-si", "--noconfirm")
}


func YayInstallPackages(pkgs ...string) error {
	if len(pkgs) == 0 {
		return nil
	}

	return Run(
		"yay",
		append([]string{"-S", "--noconfirm", "--needed"}, pkgs...)...,
	)
}
