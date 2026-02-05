package util

func SystemctlStart(service string) error {
	return Run("systemctl", "start", service)
}

func SystemctlStop(service string) error {
	return Run("systemctl", "stop", service)
}

func SystemctlRestart(service string) error {
	return Run("systemctl", "restart", service)
}

func SystemctlEnable(service string) error {
	return Run("systemctl", "enable", service)
}

func SystemctlDisable(service string) error {
	return Run("systemctl", "disable", service)
}

func EnableServiceChroot(service string) error {
	return Run(
		"sudo",
		"ln",
		"-sf",
		"/usr/lib/systemd/system/"+service,
		"/etc/systemd/system/multi-user.target.wants/"+service,
	)
}
