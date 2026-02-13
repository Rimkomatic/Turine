package install

import "turine-setup/util"

func SetShell()  error{

	if err := util.Run("chsh","-s", "/usr/bin/zsh"); err != nil {
		return err
	}
	return nil
}
