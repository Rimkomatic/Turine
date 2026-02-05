package install

import (
	"turine-setup/ui"
	"turine-setup/profiles"
)

func ChooseProfile() error {
	menu := [][]string{
		{"Server", "#ED1C24"},
		{"Carbonized", "#2ECC71"},
	}

	choice := ui.SelectMenu("Select your profile:", menu)
	ui.ClearScreen()

	switch choice {
	case 0:
		return profiles.ServerInstall()
	case 1:
		return profiles.CarbonizedInstall()
	default:
		return nil
	}
}
