package validateerrors

import "fmt"

const (
	// PermissionDenied ...
	PermissionDenied = "Permission denied"
)

// Apt ...
func Apt(err error, posError string) error {
	errStr := err.Error()
	if errStr == "exit status 100" {
		errStr = PermissionDenied
	}
	return CustomError(errStr, posError)
}

// Snap ...
func Snap(err error, posError string) error {
	errStr := err.Error()
	if errStr == "exit status 1" {
		errStr = PermissionDenied
	}
	return CustomError(errStr, posError)
}

// NPM ...
func NPM(err error, posError string) error {
	errStr := err.Error()
	if errStr == "exit status 243" {
		errStr = PermissionDenied
	}
	return CustomError(errStr, posError)
}

// Yarn ...
func Yarn(err error, posError string) error {
	errStr := err.Error()
	return CustomError(errStr, posError)
}

// PIP ...
func PIP(err error, posError string) error {
	errStr := err.Error()
	if errStr == "exit status 1" {
		errStr = PermissionDenied
	}
	return CustomError(errStr, posError)
}

// CustomError ...
func CustomError(err, posError string) error {
	return fmt.Errorf("%s [%s]", err, posError)
}
