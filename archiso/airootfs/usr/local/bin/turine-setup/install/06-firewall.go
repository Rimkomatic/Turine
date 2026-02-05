package install

import "turine-setup/util"

var FirewallDeps = []string{
	"ufw",
}

func ConfigFirewall() error {
	if err := util.PacmanInstall(FirewallDeps...); err != nil {
		return err
	}

	// sane defaults
	if err := util.Run("sudo", "ufw", "default", "deny", "incoming"); err != nil {
		return err
	}

	if err := util.Run("sudo", "ufw", "default", "allow", "outgoing"); err != nil {
		return err
	}

	if err := util.EnableServiceChroot("ufw.service"); err != nil {
		return err
	}

	if err := util.Run("sudo", "ufw", "--force", "enable"); err != nil {
		return err
	}

	return nil
}
