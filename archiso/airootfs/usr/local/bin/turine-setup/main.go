package main

import (
	"fmt"
	"log"

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

	menu := [][]string{
		{"Neovim", "#57A0D3"},
		{"Git", "#F05033"},
		{"HTop", "#00FF7F"},
	}

	choice := ui.SelectMenu("Select an app to install:", menu)

	ui.ClearScreen()

	var err error
	switch choice {
	case 0:
		err = util.PacmanInstall("neovim")
	case 1:
		err = util.PacmanInstall("git")
	case 2:
		err = util.PacmanInstall("htop")
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Installation complete.")
}
