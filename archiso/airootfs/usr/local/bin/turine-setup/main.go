package main

import (
	"fmt"

	"turine-setup/install"
	"turine-setup/ui"
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
	
	ui.ClearScreen()
	banner()


	fmt.Println("Installation complete.")
}
