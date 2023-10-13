package cmd

import (
	"fmt"
	"gostoso/functions/file_functions"
	"log"

	"github.com/spf13/cobra"
)

// expandvarsCmd represents the expandvars command
var expandvarsCmd = &cobra.Command{
	Use:   "expandvars",
	Short: "Expands (interpolate) environment variables in a file.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("expandvars called")
		// parse flags
		input, _ := cmd.Flags().GetString("input-file")
		output, _ := cmd.Flags().GetString("output-file")

		switch {
		case input == "":
			log.Fatal("You must set the flag --input-file")
		case output == "":
			log.Fatal("You must set the flag --output-file")
		default:
			fmt.Println("the common behaviour is interpolate variables into the file.")
			file_functions.ExpandVars(input, output)
		}
	},
}

func init() {
	rootCmd.AddCommand(expandvarsCmd)

	expandvarsCmd.Flags().StringP("input-file", "i", "", "Input template file with env. $VARIABLE.")
	expandvarsCmd.Flags().StringP("output-file", "o", "", "Output file name.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// expandvarsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// expandvarsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
