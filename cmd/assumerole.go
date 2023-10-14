/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gostoso/functions/aws/assumeRole"
)

// assumeroleCmd represents the assumerole command
var assumeroleCmd = &cobra.Command{
	Use:   "assumerole",
	Short: "assumes a role with the profile passed as parameter",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// parse flags
		inputProfile, _ := cmd.Flags().GetString("profile")
		inputRolearn, _ := cmd.Flags().GetString("role-arn")
		checkEval, _ := cmd.Flags().GetBool("eval")
		switch {
		case checkEval:
			assumeRole.AssumeRole(inputProfile, inputRolearn, true)
		default:
			fmt.Printf("flags checked...assuming role %s with profile %s\n", inputRolearn, inputProfile)
			assumeRole.AssumeRole(inputProfile, inputRolearn, false)
		}

	},
}

func init() {
	awsCmd.AddCommand(assumeroleCmd)

	// flags
	assumeroleCmd.Flags().StringP("profile", "p", "", "Profile whose has the hability to assume this role")
	assumeroleCmd.Flags().StringP("role-arn", "r", "", "Role that will be assumed by profile")
	assumeroleCmd.Flags().BoolP("eval", "e", false, "use \"eval $(gostoso --profile '<your profile>' --arn '<your arn>' --eval)\" command to execute assumerole, retrieve credentials and set to aws environment variables automatically")
	//required flags
	assumeroleCmd.MarkFlagsRequiredTogether("profile", "role-arn")
}
