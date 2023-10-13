package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the CLI version.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		version := "0.1.0"
		fmt.Printf("Gostoso Version: v%s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
