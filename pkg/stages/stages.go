package stages

import "fmt"

const (
	// UpgradePackages ...
	UpgradePackages = "Upgrade packages"
	// SearchUpdates ...
	SearchUpdates = "Search updates"
	// UpdatedSuccess ...
	UpdatedSuccess = "Updated Successfully"
)

// UpgradePackage ...
func UpgradePackage(pack string) string {
	return fmt.Sprintf("Upgrade package - %s", pack)
}
