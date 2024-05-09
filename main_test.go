package main_test

import (
	"testing"
	"runtime"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	binName = "gostoso"
	teste string
)

const (
	inputfile = "./testdata/expandvars/templatefile.txt"
	resultFile = "./testdata/expandvars/result.txt"
	goldenFile = "./testdata/expandvars/expected.txt"
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

	os.Exit(result)
}

func TestGostosoCliFileFunctions(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("ExpandVarsHelpMenu", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "expandvars", "--help")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
			}
})
	t.Run("ExpandVarsWriteFile", func(t *testing.T) {

		os.Setenv("TESTE", "TESTE")
		variable := os.Getenv("TESTE")
		fmt.Printf("environment variable 'TESTE' : %s\n", variable)
	
		cmd := exec.Command(cmdPath, "expandvars", "--input-file", inputfile, "--output-file", resultFile)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
			}
		
})
}