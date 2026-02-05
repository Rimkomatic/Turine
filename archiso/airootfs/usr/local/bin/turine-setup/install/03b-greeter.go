package install

import (
	"turine-setup/ui"
	"turine-setup/util"
)

func InstallGreeter() error {
	
	menu := [][]string{
		{"ly", "#ED1C24"},
		{"gdm", "#2ECC71"},
		{"cosmic", "#120871"},
	}

	choice := ui.SelectMenu("Select your profile:", menu)
	
	switch choice {
	case 0:
		return installLy() 
	case 1:
		return installGDM() 
	case 2:
		return installCosmic() 
	default:
		return nil
		
	}
}

func installLy() error {
	if err := util.PacmanInstall("ly"); err != nil {
		return err
	}

	if err := util.EnableServiceChroot("ly@tty2.service"); err != nil {
		return err
	}

	return nil
}

func installGDM() error {
	if err := util.PacmanInstall("gdm"); err != nil {
		return err
	}

	if err := util.EnableServiceChroot("gdm.service"); err != nil {
		return err
	}

	return nil
}

func installCosmic() error {
	if err := util.PacmanInstall("cosmic-greeter"); err != nil {
		return err
	}

	if err := util.EnableServiceChroot("cosmic-greeter-daemon.service"); err != nil {
		return err
	}
 
	if err := util.SystemctlEnable("cosmic-greeter.service"); err != nil {
		return err
	}

	return nil
}
