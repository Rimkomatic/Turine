package install

import (
	"os"

	"turine-setup/util"
)

func InstallYay() error {
	wd, err := os.Getwd()

	if err := util.CDHome(); err != nil {
		return err
	}

	if err := util.PacmanInstall("base-devel", "git"); err != nil {
		return err
	}

	if err != nil {
		return err
	}
	defer os.Chdir(wd)

	if err := util.GitClone("https://aur.archlinux.org/yay.git"); err != nil {
		return err
	}

	if err := util.ChangeDirectory("yay"); err != nil {
		return err
	}

	if err := util.MakepkgInstall(); err != nil {
		return err
	}

	return nil
}
