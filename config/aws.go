package config //Specific configurations found in config.yaml

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

const (
	AWSRegion      = "us-west-2" // Replace with your desired region
	AWSCredentials = "default"   // Replace with your desired credential profile
)

// GetAWSCredentials returns the AWS credentials based on the specified profile
func GetAWSCredentials() *credentials.Credentials {
	creds := credentials.NewSharedCredentials("", AWSCredentials)
	return creds
}
