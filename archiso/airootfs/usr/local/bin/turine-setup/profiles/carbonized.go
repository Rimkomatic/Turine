package profiles

import "turine-setup/util"

func CarbonizedInstall()  error {
	if err := util.YayInstallPackages(CarbonizedDepsYay...); err != nil {
		return err
	}
	
	if err := util.PacmanInstall(CarbonizedDeps...); err != nil {
		return err
	}

	return nil
}
