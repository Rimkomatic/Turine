package install

import "turine-setup/util"

var AudioDeps = []string{
	"pipewire",
	"pipewire-alsa",
	"pipewire-jack",
	"pipewire-pulse",
	"gst-plugin-pipewire",
	"libpulse",
	"wireplumber",
}

func ConfigAudio() error {
	if err := util.PacmanInstall(AudioDeps...); err != nil {
		return err
	}

	if err := util.RunAsUser(
		"systemctl", "--user", "enable", "--now", "pipewire.service",
	); err != nil {
		return err
	}

	if err := util.RunAsUser(
		"systemctl", "--user", "enable", "--now", "pipewire-pulse.service",
	); err != nil {
		return err
	}

	if err := util.RunAsUser(
		"systemctl", "--user", "enable", "--now", "wireplumber.service",
	); err != nil {
		return err
	}

	return nil
}

