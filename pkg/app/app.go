package app

import (
	"fmt"
	"sync"
	"time"
	"update-packages/pkg/packagemanagers/apt"
	"update-packages/pkg/packagemanagers/npm"
	"update-packages/pkg/packagemanagers/pip"
	"update-packages/pkg/packagemanagers/snap"
	"update-packages/pkg/packagemanagers/yarn"
	"update-packages/pkg/stages"
	"update-packages/pkg/vars"
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
	go UpdateApt(&wg)
	wg.Add(1)
	go UpdateSnap(&wg)
	wg.Add(1)
	go UpdateNPM(&wg)
	wg.Add(1)
	go UpdateYarn(&wg)
	wg.Add(1)
	go UpdatePip(&wg)

	wg.Wait()

	if len(errorsList) > 0 {
		ShowErrors(errorsList)
	} else {
		fmt.Println("All Packages Updated Successfully")
	}

	endTime := time.Now().UnixNano()
	duration := float64(time.Duration(endTime-startTime)) / float64(time.Second)
	fmt.Printf("App finished in %f seconds\n", duration)
}

// UpdateApt ...
func UpdateApt(wg *sync.WaitGroup) {
	defer wg.Done()
	if err := apt.Run(); err != nil {
		errorsList = append(errorsList, CustomError(err, vars.APT))
		return
	}
	fmt.Println(CustomSuccess(vars.APT))
}

// UpdateSnap ...
func UpdateSnap(wg *sync.WaitGroup) {
	defer wg.Done()
	if err := snap.Run(); err != nil {
		errorsList = append(errorsList, CustomError(err, vars.Snap))
		return
	}
	fmt.Println(CustomSuccess(vars.Snap))
}

// UpdateNPM ...
func UpdateNPM(wg *sync.WaitGroup) {
	defer wg.Done()
	if err := npm.Run(); err != nil {
		errorsList = append(errorsList, CustomError(err, vars.NPM))
		return
	}
	fmt.Println(CustomSuccess(vars.NPM))
}

// UpdateYarn ...
func UpdateYarn(wg *sync.WaitGroup) {
	defer wg.Done()
	if err := yarn.Run(); err != nil {
		errorsList = append(errorsList, CustomError(err, vars.Yarn))
		return
	}
	fmt.Println(CustomSuccess(vars.Yarn))
}

// UpdatePip ...
func UpdatePip(wg *sync.WaitGroup) {
	defer wg.Done()
	if err := pip.Run(); err != nil {
		errorsList = append(errorsList, CustomError(err, vars.PIP))
		return
	}
	fmt.Println(CustomSuccess(vars.PIP))
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
