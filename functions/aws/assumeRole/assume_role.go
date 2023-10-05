package assumeRole

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

func AssumeRole() {
	// Initial credentials loaded from SDK's default credential chain. Such as
	// the environment, shared credentials (~/.aws/credentials), or EC2 Instance
	// Role. These credentials will be used to to make the STS Assume Role API.
	sess := session.Must(session.NewSession())
	sess.Config.Credentials.Get()
	// Create the credentials from AssumeRoleProvider to assume the role
	// referenced by the "myRoleARN" ARN.
	creds := stscreds.NewCredentials(sess, "arn:aws:iam::622734844733:role/assumeRoleS3All")
	// Create service client value configured for credentials
	// from assumed role.
	fmt.Printf("Credentials: ", creds)
}
