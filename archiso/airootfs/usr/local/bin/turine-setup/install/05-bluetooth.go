package install

import "turine-setup/util"

var BluetoothDeps = []string{
	"bluez",
	"bluez-utils",
}

func ConfigBluetooth() error {
	if err := util.PacmanInstall(BluetoothDeps...); err != nil {
		return err
	}

	if err := util.EnableServiceChroot("bluetooth.service"); err != nil {
		return err
	}

	return nil
}
