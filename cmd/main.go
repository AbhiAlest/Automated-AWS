package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AbhiAlest/Automated-AWS/pkg/ec2"
	"github.com/AbhiAlest/Automated-AWS/pkg/elbv2"
	"github.com/AbhiAlest/Automated-AWS/pkg/autoscaling"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func main() {
	// Call the functions to create AWS resources
	ec2.CreateEC2Instance()
	elbv2.CreateLoadBalancer()
	autoscaling.CreateAutoScalingGroup()

	// Additional function calls or logic
	createVPC()
	createS3Bucket()
	launchEC2Instances(2)
	createALB()
}

func createVPC() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create VPC
	ec2Client := ec2.New(sess)
	vpcInput := &ec2.CreateVpcInput{
		CidrBlock: aws.String("10.0.0.0/16"),
	}
	vpcOutput, err := ec2Client.CreateVpc(vpcInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("VPC created:", *vpcOutput.Vpc.VpcId)

	// Enable DNS support in the VPC
	enableDNSInput := &ec2.EnableVpcClassicLinkDnsSupportInput{
		VpcId: vpcOutput.Vpc.VpcId,
	}
	_, err = ec2Client.EnableVpcClassicLinkDnsSupport(enableDNSInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DNS support enabled for VPC:", *vpcOutput.Vpc.VpcId)
}

func createS3Bucket() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create S3 bucket
	s3Client := s3.New(sess)
	bucketInput := &s3.CreateBucketInput{
		Bucket: aws.String("my-bucket"),
	}
	_, err = s3Client.CreateBucket(bucketInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("S3 bucket created: my-bucket")

	// Upload a file to the S3 bucket
	uploadFileInput := &s3manager.UploadInput{
		Bucket: aws.String("my-bucket"),
		Key:    aws.String("my-file.txt"),
		Body:   strings.NewReader("Hello, World!"),
	}
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(uploadFileInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File uploaded to S3 bucket")
}

func launchEC2Instances(numInstances int) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Launch EC2 instances
	ec2Client := ec2.New(sess)
	runInstancesInput := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-12345678"),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(int64(numInstances)),
		MaxCount:     aws.Int64(int64(numInstances)),
	}
	_, err = ec2Client.RunInstances(runInstancesInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Launched %d EC2 instances\n", numInstances)
}

func createALB() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create ELBv2
	elbv2Client := elbv2.New(sess)
	createLBInput := &elbv2.CreateLoadBalancerInput{
		Name: aws.String("my-load-balancer"),
		Subnets: []*string{
			aws.String("subnet-12345678"),
			aws.String("subnet-98765432"),
		},
		SecurityGroups: []*string{
			aws.String("sg-12345678"),
		},
	}
	_, err = elbv2Client.CreateLoadBalancer(createLBInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ELBv2 created")

	// Create a target group
	createTGInput := &elbv2.CreateTargetGroupInput{
		Name:     aws.String("my-target-group"),
		Protocol: aws.String("HTTP"),
		Port:     aws.Int64(80),
		VpcId:    aws.String("vpc-12345678"),
	}
	_, err = elbv2Client.CreateTargetGroup(createTGInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Target group created")

	// Attach instances to the target group
	attachInstancesInput := &elbv2.RegisterTargetsInput{
		TargetGroupArn: aws.String("arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/my-target-group/abcdef123456"),
		Targets: []*elbv2.TargetDescription{
			{
				Id: aws.String("i-1234567890abcdef0"),
			},
			{
				Id: aws.String("i-0987654321fedcba0"),
			},
		},
	}
	_, err = elbv2Client.RegisterTargets(attachInstancesInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Instances attached to the target group")
}
