package autoscaling

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/AbhiAlest/Automated-AWS/config"
)

// CreateAutoScalingGroup creates an Auto Scaling group with the given configuration
func CreateAutoScalingGroup() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWSRegion),
	})
	if err != nil {
		panic(err)
	}

	svc := autoscaling.New(sess)

	// Input parameters for creating an Auto Scaling group
	input := &autoscaling.CreateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String("MyAutoScalingGroup"), // Replace with the desired name
		LaunchTemplate: &autoscaling.LaunchTemplateSpecification{
			LaunchTemplateName: aws.String("MyLaunchTemplate"), // Replace with the desired launch template name
			Version:            aws.String("$Latest"),
		},
		MinSize:          aws.Int64(1),
		MaxSize:          aws.Int64(5),
		DesiredCapacity:  aws.Int64(2),
		VPCZoneIdentifier: aws.String("subnet-12345678"), // Replace with the desired subnet ID
		Tags: []*autoscaling.Tag{
			{
				Key:               aws.String("Environment"),
				Value:             aws.String("Production"),
				PropagateAtLaunch: aws.Bool(true),
			},
			{
				Key:               aws.String("Project"),
				Value:             aws.String("MyProject"),
				PropagateAtLaunch: aws.Bool(true),
			},
		},
	}

	result, err := svc.CreateAutoScalingGroup(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Auto Scaling group created:", *result.AutoScalingGroupARN)

	// Suspend Auto Scaling Processes
	suspendInput := &autoscaling.ScalingProcessQuery{
		AutoScalingGroupName: aws.String("MyAutoScalingGroup"), // Replace with the name of your Auto Scaling group
		ScalingProcesses: []*string{
			aws.String("AlarmNotification"),
		},
	}

	_, err = svc.SuspendProcesses(suspendInput)
	if err != nil {
		panic(err)
	}

	fmt.Println("Auto Scaling processes suspended")

	// Resume Auto Scaling Processes
	resumeInput := &autoscaling.ScalingProcessQuery{
		AutoScalingGroupName: aws.String("MyAutoScalingGroup"), // Replace with the name of your Auto Scaling group
		ScalingProcesses: []*string{
			aws.String("AlarmNotification"),
		},
	}

	_, err = svc.ResumeProcesses(resumeInput)
	if err != nil {
		panic(err)
	}

	fmt.Println("Auto Scaling processes resumed")

	// Attach Load Balancer to Auto Scaling Group
	attachInput := &autoscaling.AttachLoadBalancerTargetGroupsInput{
		AutoScalingGroupName: aws.String("MyAutoScalingGroup"), // Replace with the name of your Auto Scaling group
		TargetGroupARNs: []*string{
			aws.String("arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/my-target-group/abcdef123456"),
		},
	}

	_, err = svc.AttachLoadBalancerTargetGroups(attachInput)
	if err != nil {
		panic(err)
	}

	fmt.Println("Load balancer attached to Auto Scaling group")
}

