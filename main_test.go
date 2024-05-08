package main_test

import (
	"testing"
	"runtime"
	"fmt"
	"os"
	"os/exec"
)

var (
	binName = "gostoso"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...if GOOS=windows it will generate a .exe file")
//define operation system
	if runtime.GOOS == "windows" {
		binName += ".exe"
		fmt.Printf("Generating %s ", binName)
	}
// build the tool depending on OS type
	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...removing binaries")
	os.Remove(binName)

	os.Exit(result)
}

