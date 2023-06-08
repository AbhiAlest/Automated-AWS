package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/AbhiAlest/Automated-AWS/config"
)

func CreateEC2Instance() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWSRegion),
	})
	if err != nil {
		panic(err)
	}

	svc := ec2.New(sess)

	// Implement your logic to create an EC2 instance
	// using the svc object and the necessary input parameters
}

