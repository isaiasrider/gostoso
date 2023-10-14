/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gostoso/functions/aws/assumeRole"
	"log"

	"github.com/spf13/cobra"
)

// assumeroleCmd represents the assumerole command
var assumeroleCmd = &cobra.Command{
	Use:   "assumerole",
	Short: "assumes a role with the profile passed as parameter",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("assumerole command called \n")
		// parse flags
		inputProfile, _ := cmd.Flags().GetString("profile")
		inputRolearn, _ := cmd.Flags().GetString("role-arn")

		switch {
		case inputProfile == "":
			log.Fatal("You must set the flag --profile")
		case inputRolearn == "":
			log.Fatal("You must set the flag --role-arn")
		default:
			fmt.Printf("flags checked...assuming role %s with profile %s\n", inputRolearn, inputProfile)
			assumeRole.AssumeRole(inputProfile, inputRolearn)
		}

	},
}

func init() {
	awsCmd.AddCommand(assumeroleCmd)

	assumeroleCmd.Flags().StringP("profile", "p", "", "Profile whose has the hability to assume this role")
	assumeroleCmd.Flags().StringP("role-arn", "r", "", "Role that will be assumed by profile")
}
