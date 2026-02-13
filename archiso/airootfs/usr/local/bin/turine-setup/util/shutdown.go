package util

import (
	"os/exec"
	"time"
)

func ShutdownAfter5Seconds() error {
	time.Sleep(5 * time.Second)
	return exec.Command("shutdown", "-h", "now").Run()
}
