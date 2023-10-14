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

func AssumeRole(profile string, rolearn string, eval bool) {
	// load a aws profile passed as flag

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	if err != nil {
		panic(err)
	}

	// Create the credentials from AssumeRoleProvider to assume the role
	// referenced by the "rolearn" flag.
	stsClient := sts.NewFromConfig(cfg)
	creds := stscreds.NewAssumeRoleProvider(stsClient, rolearn)

	credentials, _ := creds.Retrieve(context.TODO())
	if eval {
		awsAccessKeyId := credentials.AccessKeyID
		awsSecretAccessKey := credentials.SecretAccessKey
		awsSessionToken := credentials.SessionToken
		fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", awsAccessKeyId)
		fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", awsSecretAccessKey)
		fmt.Printf("export AWS_SESSION_TOKEN=%s\n", awsSessionToken)
	} else {
		fmt.Printf("Loading Profile: %s\n\n", profile)
		fmt.Printf("Your AccessKeyId from AssumedRole: %s\n\n", credentials.AccessKeyID)
		fmt.Printf("Your SecretAccessKey from AssumedRole: %s\n\n", credentials.SecretAccessKey)
		fmt.Printf("Your SessionToken: %s\n\n", credentials.SessionToken)
		fmt.Printf("Your Credentials Expires: %s\n", credentials.Expires)
		fmt.Println()
		identity, err := stsClient.GetCallerIdentity(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Account: %s, Arn: %s\n", aws.ToString(identity.Account), aws.ToString(identity.Arn))
	}
}

// call sts get-caller-identity to confirm the user
