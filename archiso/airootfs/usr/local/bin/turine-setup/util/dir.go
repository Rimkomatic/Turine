package util

import "os" 

func ChangeDirectory(path string) error {
	return os.Chdir(path)
}

func CDHome() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return os.Chdir(home)
}
