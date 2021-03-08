package npm

import (
	"os/exec"
	"update-packages/pkg/stages"
	"update-packages/pkg/validateerrors"
)

// Run ...
func Run() error {
	if err := UpgradePackages(); err != nil {
		return validateerrors.ValidateErrors(err, stages.UpgradePackages)
	}

	return nil
}

// UpgradePackages ...
func UpgradePackages() error {
	return exec.Command("npm", "-g", "upgrade").Run()
}
