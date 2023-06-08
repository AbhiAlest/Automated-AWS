package elbv2

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/AbhiAlest/Automated-AWS/config"
)

func CreateLoadBalancer() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWSRegion),
	})
	if err != nil {
		panic(err)
	}

	svc := elbv2.New(sess)

	// Input parameters for creating a load balancer
	input := &elbv2.CreateLoadBalancerInput{
		Name:        aws.String("MyLoadBalancer"), // Replace with the desired name
		Subnets:     []*string{aws.String("subnet-12345678")}, // Replace with the desired subnet ID
		SecurityGroups: []*string{aws.String("sg-12345678")}, // Replace with the desired security group ID
	}

	// Optionally, configure the load balancer listeners
	listeners := []*elbv2.Listener{
		{
			Port:     aws.Int64(80),
			Protocol: aws.String("HTTP"),
			DefaultActions: []*elbv2.Action{
				{
					Type:           aws.String("forward"),
					TargetGroupArn: aws.String("arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/my-target-group/abcdef123456"),
				},
			},
		},
	}

	input.SetListeners(listeners)

	// Create the load balancer
	result, err := svc.CreateLoadBalancer(input)
	if err != nil {
		panic(err)
	}

	// Get the ARN of the created load balancer
	loadBalancerARN := *result.LoadBalancers[0].LoadBalancerArn
	fmt.Println("Load balancer created:", loadBalancerARN)

	// Modify the load balancer attributes
	modifyInput := &elbv2.ModifyLoadBalancerAttributesInput{
		LoadBalancerArn: aws.String(loadBalancerARN),
		Attributes: []*elbv2.LoadBalancerAttribute{
			{
				Key:   aws.String("access_logs.s3.enabled"),
				Value: aws.String("true"),
			},
			{
				Key:   aws.String("access_logs.s3.bucket"),
				Value: aws.String("my-log-bucket"),
			},
			{
				Key:   aws.String("idle_timeout.timeout_seconds"),
				Value: aws.String("300"),
			},
		},
	}

	_, err = svc.ModifyLoadBalancerAttributes(modifyInput)
	if err != nil {
		panic(err)
	}

	fmt.Println("Load balancer modified with attributes")

	// Delete the load balancer
	deleteInput := &elbv2.DeleteLoadBalancerInput{
		LoadBalancerArn: aws.String(loadBalancerARN),
	}

	_, err = svc.DeleteLoadBalancer(deleteInput)
	if err != nil {
		panic(err)
	}

	fmt.Println("Load balancer deleted:", loadBalancerARN)
}

