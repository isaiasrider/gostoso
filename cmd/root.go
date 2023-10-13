package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gostoso",
	Short: "Ghost Open-Source Tools for System Operations",
	Long:  `GOSToSO: a CLI for deployment automation.`,
	// Uncomment the following line if your bare application has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gostoso.yaml)")

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose mode (not implemented yet).")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug mode (not implemented yet).")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
