/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gostoso/functions/file_functions"
	"log"
)

// expandvarsCmd represents the expandvars command
var expandvarsCmd = &cobra.Command{
	Use:   "expandvars",
	Short: "Expands variables in a file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("expandvars called")
		// parse flags
		input, _ := cmd.Flags().GetString("inputfile")
		output, _ := cmd.Flags().GetString("outputfile")

		switch {
		case input == "":
			log.Fatal("you need to feed the flag --inputfile")

		case output == "":
			log.Fatal("you need to feed the flag --outputfile")
		default:
			fmt.Println("the commong behaviour is pop-up variables into the file")
			file_functions.ExpandVars(input, output)

		}
	},
}

func init() {
	rootCmd.AddCommand(expandvarsCmd)

	expandvarsCmd.Flags().StringP("inputfile", "i", "", "Input template file with $VARIABLES")
	expandvarsCmd.Flags().StringP("outputfile", "o", "", "Output file with any name you want")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// expandvarsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// expandvarsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
