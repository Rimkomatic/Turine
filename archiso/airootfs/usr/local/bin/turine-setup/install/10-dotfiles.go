package install

import "turine-setup/util"

func InstallDotfiles()  error{

	if err := util.PacmanInstall("stow"); err != nil {
		return err
	}

	if err := util.CDHome(); err != nil {
		return err
	}
	
	if err := util.GitClone("https://github.com/Rimkomatic/dotfiles.git"); err != nil {
		return err
	}

	if err := util.ChangeDirectory("dotfiles"); err != nil {
		return err
	}

	if err := util.StowDotfiles("fastfetch hyprland matugen niri nvim rofi starship tmux waypaper wlogout yazi zathura zsh env-eventhorizon"); err != nil {
		return err
	}

	return nil
}
