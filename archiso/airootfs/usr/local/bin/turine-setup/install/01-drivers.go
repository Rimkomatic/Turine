package install

import (
	"turine-setup/ui"
	"turine-setup/util"
)

func CpuDrivers() error {
	menu := [][]string{
		{"AMD (amd-ucode)", "#ED1C24"},
		{"Intel (intel-ucode)", "#0071C5"},
	}

	choice := ui.SelectMenu("Select your cpu:", menu)
	ui.ClearScreen()

	switch choice {
	case 0:
		return util.PacmanInstall("amd-ucode")
	case 1:
		return util.PacmanInstall("intel-ucode")
	default:
		return nil
	}
}

func GpuDrivers() error {
	menu := [][]string{
		{"AMD (mesa / amdgpu)", "#ED1C24"},
		{"Intel (mesa / i915)", "#0071C5"},
		{"NVIDIA (no CUDA)", "#2ECC71"},
		{"NVIDIA (with CUDA)", "#2ECC71"},
	}

	choice := ui.SelectMenu("Select your GPU:", menu)
	ui.ClearScreen()

	switch choice {

	case 0:
		return util.YayInstallPackages(
			"mesa",
			"vulkan-radeon",
			"libva-mesa-driver",
		)
	case 1:
		return util.YayInstallPackages(
			"mesa",
			"vulkan-intel",
		)
	case 2:
		return util.PacmanInstall(
			"nvidia-open",
			"nvidia-utils",
			"nvidia-settings",
			"opencl-nvidia",
		)
	case 3:
		if err := util.PacmanInstall(
			"nvidia-open",
			"nvidia-utils",
			"nvidia-settings",
			"opencl-nvidia",
			"nvidia-container-toolkit",
		); err != nil {
			return err
		}
		if err := util.PacmanInstallU(
			"https://archive.archlinux.org/packages/g/gcc/gcc-14.2.1+r134+gab884fffe3fc-2-x86_64.pkg.tar.zst",
			"https://archive.archlinux.org/packages/g/gcc-libs/gcc-libs-14.2.1+r134+gab884fffe3fc-2-x86_64.pkg.tar.zst",
			"https://archive.archlinux.org/packages/b/binutils/binutils-2.42-1-x86_64.pkg.tar.zst",
		); err != nil {
			return err
		}
		return util.PacmanInstallU(
			"https://archive.archlinux.org/packages/c/cuda/cuda-12.8.1-3-x86_64.pkg.tar.zst",
			"https://archive.archlinux.org/packages/c/cuda-tools/cuda-tools-12.8.1-3-x86_64.pkg.tar.zst",
			"https://archive.archlinux.org/packages/c/cudnn/cudnn-9.7.0.66-1-x86_64.pkg.tar.zst",
		)

	default:
		return nil
	}
}
