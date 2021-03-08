package app

import (
	"fmt"
	"sync"
	"time"
	"update-packages/pkg/apt"
	"update-packages/pkg/npm"
	"update-packages/pkg/snap"
	"update-packages/pkg/stages"
	"update-packages/pkg/vars"
	"update-packages/pkg/yarn"
)

var (
	errorsList  []error
	successList []string
)

// Run ...
func Run() {
	startTime := time.Now().UnixNano()

	var wg sync.WaitGroup

	wg.Add(1)
	UpdateApt(&wg)
	wg.Add(1)
	UpdateSnap(&wg)
	wg.Add(1)
	UpdateNPM(&wg)
	wg.Add(1)
	UpdateYarn(&wg)

	wg.Wait()

	if len(successList) > 0 {
		ShowSuccess(successList)
	}
	if len(errorsList) > 0 {
		ShowErrors(errorsList)
	} else {
		fmt.Println("Packages Updated Successfully")
	}

	endTime := time.Now().UnixNano()
	fmt.Println(endTime - startTime)
}

// UpdateApt ...
func UpdateApt(wg *sync.WaitGroup) {
	defer wg.Done()
	if err := apt.Run(); err != nil {
		errorsList = append(errorsList, CustomError(err, vars.APT))
		return
	}
	successList = append(successList, CustomSuccess(vars.APT))
}

// UpdateSnap ...
func UpdateSnap(wg *sync.WaitGroup) {
	defer wg.Done()
	if err := snap.Run(); err != nil {
		errorsList = append(errorsList, CustomError(err, vars.Snap))
		return
	}
	successList = append(successList, CustomSuccess(vars.Snap))
}

// UpdateNPM ...
func UpdateNPM(wg *sync.WaitGroup) {
	defer wg.Done()
	if err := npm.Run(); err != nil {
		errorsList = append(errorsList, CustomError(err, vars.NPM))
		return
	}
	successList = append(successList, CustomSuccess(vars.NPM))
}

// UpdateYarn ...
func UpdateYarn(wg *sync.WaitGroup) {
	defer wg.Done()
	if err := yarn.Run(); err != nil {
		errorsList = append(errorsList, CustomError(err, vars.Yarn))
		return
	}
	successList = append(successList, CustomSuccess(vars.Yarn))
}

// CustomError ...
func CustomError(err error, packageManager string) error {
	return fmt.Errorf("%s [%s]", err, packageManager)
}

// CustomSuccess ...
func CustomSuccess(packageManager string) string {
	return fmt.Sprintf("%s [%s]", stages.UpdatedSuccess, packageManager)
}

// ShowErrors ...
func ShowErrors(errs []error) {
	for _, err := range errorsList {
		fmt.Println(err)
	}
}

// ShowSuccess ...
func ShowSuccess(listSuccess []string) {
	for _, mess := range listSuccess {
		fmt.Println(mess)
	}
}
