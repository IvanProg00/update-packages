package apt

import (
	"os/exec"
	"update-packages/pkg/stages"
	"update-packages/pkg/validateerrors"
)

// Run ...
func Run() error {
	if err := Update(); err != nil {
		return validateerrors.Apt(err, stages.SearchUpdates)
	}

	if err := UpgradePackages(); err != nil {
		return validateerrors.Apt(err, stages.UpgradePackages)
	}

	return nil
}

// Update ...
func Update() error {
	return exec.Command("apt", "update").Run()
}

// UpgradePackages ...
func UpgradePackages() error {
	return exec.Command("apt", "full-upgrade", "-y").Run()
}
