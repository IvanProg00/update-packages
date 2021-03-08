package validateerrors

import "fmt"

const (
	// PermissionDenied ...
	PermissionDenied = "Permission denied"
)

// ValidateErrors ...
func ValidateErrors(err error, posError string) error {
	errStr := err.Error()
	if err.Error() == "exit status 1" || err.Error() == "exit status 100" || err.Error() == "exit status 243" {
		errStr = PermissionDenied
	}
	return fmt.Errorf("%s [%s]", errStr, posError)
}
