package file_functions

import (
	"fmt"
	"log"
	"os"
)

func ExpandVars(inputfile string, outputfile string) {
	fmt.Printf("Input file: '%s'\nOutput file: '%s'", inputfile, outputfile)

	templateFile, err := os.ReadFile(inputfile)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(outputfile, []byte(os.ExpandEnv(string(templateFile))), 0777)
}
