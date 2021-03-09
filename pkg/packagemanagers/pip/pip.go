package pip

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"update-packages/pkg/stages"
	"update-packages/pkg/validateerrors"
)

// Run ...
func Run() error {
	updates, err := ShowUpdates()
	if err != nil {
		return validateerrors.PIP(err, stages.SearchUpdates)
	}
	packages := FilterPackages(updates)

	if err := UpgradePackages(packages); err != nil {
		return validateerrors.PIP(err, stages.UpgradePackages)
	}

	return nil
}

// ShowUpdates ...
func ShowUpdates() (string, error) {
	data, err := exec.Command("pip", "list", "--outdated").Output()
	return string(data), err
}

// FilterPackages ...
func FilterPackages(updates string) []string {
	packages := strings.Split(updates, "\n")
	packages = packages[2 : len(packages)-1]

	for i, pack := range packages {
		packages[i] = pack[:strings.Index(pack, " ")]
	}

	return packages
}

// UpgradePackages ...
func UpgradePackages(packages []string) error {
	var errSer error
	var wg sync.WaitGroup

	for _, pack := range packages {
		wg.Add(1)
		go func(pack string, wg *sync.WaitGroup) {
			defer wg.Done()
			err := exec.Command("pip", "install", "--upgrade", pack).Run()
			if err != nil {
				fmt.Println(validateerrors.PIP(err, stages.UpgradePackage(pack)))
				errSer = err
			}
		}(pack, &wg)
	}

	wg.Wait()
	return errSer
}
