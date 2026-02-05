package profiles

import "turine-setup/util"

func ServerInstall()  error {
	return util.YayInstallPackages(ServerDeps...)
}
