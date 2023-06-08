package ec2

import (
	"fmt"
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

	// Input parameters for creating an EC2 instance
	input := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-12345678"),     // Replace with the desired AMI ID
		InstanceType: aws.String("t2.micro"),         // Replace with the desired instance type
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
		SecurityGroupIds: []*string{
			aws.String("sg-12345678"),                // Replace with the desired security group ID
		},
		SubnetId: aws.String("subnet-12345678"),      // Replace with the desired subnet ID
	}

	// Optionally, add tags to the EC2 instance
	tags := []*ec2.Tag{
		{
			Key:   aws.String("Name"),
			Value: aws.String("MyInstance"),
		},
		{
			Key:   aws.String("Environment"),
			Value: aws.String("Production"),
		},
	}

	input.SetTagSpecifications([]*ec2.TagSpecification{
		{
			ResourceType: aws.String("instance"),
			Tags:         tags,
		},
	})

	// Create the EC2 instance
	result, err := svc.RunInstances(input)
	if err != nil {
		panic(err)
	}

	// Get the instance ID of the created instance
	instanceID := *result.Instances[0].InstanceId
	fmt.Println("EC2 instance created:", instanceID)

	// Modify the EC2 instance
	modifyInput := &ec2.ModifyInstanceAttributeInput{
		InstanceId: aws.String(instanceID),
		Attributes: []*ec2.InstanceAttributeName{
			aws.String("userData"),
		},
		UserData: &ec2.BlobAttributeValue{
			Value: []byte("#!/bin/bash\necho 'Hello, Instance!'"),
		},
	}

	_, err = svc.ModifyInstanceAttribute(modifyInput)
	if err != nil {
		panic(err)
	}

	fmt.Println("EC2 instance modified with user data")

	// Terminate the EC2 instance
	terminateInput := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}

	_, err = svc.TerminateInstances(terminateInput)
	if err != nil {
		panic(err)
	}

	fmt.Println("EC2 instance terminated:", instanceID)
}

