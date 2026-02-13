package util

import (
	"fmt"
	"os"
)

func FixSbinLayout() error {
	dirs := []string{"/sbin", "/usr/sbin"}

	for _, dir := range dirs {
		info, err := os.Lstat(dir)

		if os.IsNotExist(err) {
			if err := os.Symlink("/usr/bin", dir); err != nil {
				return fmt.Errorf("failed to create symlink %s: %w", dir, err)
			}
			continue
		}

		if err != nil {
			return fmt.Errorf("error stating %s: %w", dir, err)
		}

		if info.Mode()&os.ModeSymlink != 0 {
			continue
		}

		backup := dir + ".old"
		if err := os.Rename(dir, backup); err != nil {
			return fmt.Errorf("failed to backup %s: %w", dir, err)
		}

		if err := os.Symlink("/usr/bin", dir); err != nil {
			return fmt.Errorf("failed to create symlink %s: %w", dir, err)
		}
	}

	return nil
}
