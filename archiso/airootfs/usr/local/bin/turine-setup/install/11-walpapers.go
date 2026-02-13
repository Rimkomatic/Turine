package install

import "turine-setup/util"


func GetWalpapers()  error {

	if err := util.CDHome(); err != nil {
		return err
	}

	if err := util.GitCloneOn("https://github.com/Rimkomatic/Wallpapers_turine.git", "./Pictures/Wallpapers"); err != nil {
		return err
	}

	return nil
}
