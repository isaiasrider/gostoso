package assumeRole

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"log"
)

type awsCredetials aws.Credentials

func AssumeRole() {
	// load a aws profile passed as flag

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loading Profile: \n\n")

	// Create the credentials from AssumeRoleProvider to assume the role
	// referenced by the "rolearn" flag.
	stsClient := sts.NewFromConfig(cfg)
	creds := stscreds.NewAssumeRoleProvider(stsClient, "arn:aws:iam::622734844733:role/assumeRoleS3All")

	credentials, _ := creds.Retrieve(context.TODO())

	fmt.Printf("Your AccessKeyId from AssumedRole: %s\n", credentials.AccessKeyID)
	fmt.Printf("Your SecretAccessKey from AssumedRole: %s\n", credentials.SecretAccessKey)
	fmt.Printf("Your SessionToken: %s\n", credentials.SessionToken)
	fmt.Println()

	// call sts get-caller-identity to confirm the user
	identity, err := stsClient.GetCallerIdentity(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account: %s, Arn: %s\n", aws.ToString(identity.Account), aws.ToString(identity.Arn))

}
