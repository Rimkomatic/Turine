package main

import (
	"fmt"

	"turine-setup/install"
	"turine-setup/ui"
	"turine-setup/util"
)

func banner() {
	fmt.Print(`
████████╗██╗   ██╗██████╗ ██╗███╗   ██╗███████╗
╚══██╔══╝██║   ██║██╔══██╗██║████╗  ██║██╔════╝
   ██║   ██║   ██║██████╔╝██║██╔██╗ ██║█████╗  
   ██║   ██║   ██║██╔══██╗██║██║╚██╗██║██╔══╝  
   ██║   ╚██████╔╝██║  ██║██║██║ ╚████║███████╗
   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝╚══════╝

`)
}

func main() {
	ui.ClearScreen()
	banner()

	if err := install.InstallYay(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}

	ui.ClearScreen()
	banner()

	if err := install.CpuDrivers(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}

	ui.ClearScreen()
	banner()

	if err := install.GpuDrivers(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}

	ui.ClearScreen()
	banner()
	if err := install.ConfigAudio(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}

	ui.ClearScreen()
	banner()
	if err := install.ConfigBluetooth(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}

	ui.ClearScreen()
	banner()

	if err := install.ChooseProfile(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}

	ui.ClearScreen()
	banner()
	if err := install.InstallGreeter(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}
	ui.ClearScreen()
	banner()

	if err := install.ConfigFirewall(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}

	ui.ClearScreen()
	banner()

	if err := install.InstallDotfiles(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}

	if err := install.GetWalpapers(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}
	ui.ClearScreen()
	banner()

	if err := util.WhichPackage("zsh"); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}

	if err := util.FixSbinLayout(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}

	if err := install.SetShell(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}

	fmt.Println("Installation complete.")

	if err := util.ShutdownAfter5Seconds(); err != nil {
		ui.Error("Installation failed")
		ui.Error(err.Error())
		return
	}
}
