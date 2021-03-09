package yarn

import (
	"os/exec"
	"update-packages/pkg/stages"
	"update-packages/pkg/validateerrors"
)

// Run ...
func Run() error {
	if err := UpgradePackages(); err != nil {
		return validateerrors.Yarn(err, stages.UpgradePackages)
	}

	return nil
}

// UpgradePackages ...
func UpgradePackages() error {
	return exec.Command("yarn", "global", "upgrade").Run()
}
