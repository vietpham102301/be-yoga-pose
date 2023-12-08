package python

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func GetABSPath() {
	absolutePath, err := getAbsolutePath()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Absolute Path:", absolutePath)
}

func getAbsolutePath() (string, error) {
	// Use runtime.Caller to get the caller's file path
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("unable to get the caller's file path")
	}

	// Use filepath.Abs to get the absolute path
	absolutePath, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}

	return absolutePath, nil
}
