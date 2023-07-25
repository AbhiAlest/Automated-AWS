package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"github.com/spf13

	"github.com/AbhiAlest/Automated-AWS/pkg/ec2"
	"github.com/AbhiAlest/Automated-AWS/pkg/elbv2"
	"github.com/AbhiAlest/Automated-AWS/pkg/autoscaling"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

var (
	logger *log.Logger
)

func init() {
	// Initialize the logger
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
}


func main() {
	// cobra

	var rootCmd = &cobra.Command{
		Use:   "automated-aws",
		Short: "Automated-AWS is an open-source Go package that automates provisioning/management of AWS resources (ec2, elbv2, and autoscaling).",
	}

	var createEC2Cmd = &cobra.Command{
		Use:   "create-ec2",
		Short: "Create an EC2 instance",
		Run: func(cmd *cobra.Command, args []string) {
			createEC2Instance()
		},
	}

	rootCmd.AddCommand(createEC2Cmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	
	
	// Define command-line flags, see more in README
	createVPCFlag := flag.Bool("create-vpc", false, "Create VPC")
	createS3BucketFlag := flag.Bool("create-s3-bucket", false, "Create S3 bucket")
	launchEC2InstancesFlag := flag.Int("launch-ec2-instances", 0, "Launch EC2 instances")
	createALBFlag := flag.Bool("create-alb", false, "Create ALB")
	flag.Parse()

	// Create a WaitGroup to wait for all Goroutines to finish
		var wg sync.WaitGroup

	// Call the functions based on the provided flags
	if *createVPCFlag {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := createVPC(); err != nil {
				logger.Println("Error creating VPC:", err)
			}
		}()
	}
	if *createS3BucketFlag {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := createS3Bucket(); err != nil {
				logger.Println("Error creating S3 bucket:", err)
			}
		}()
	}
	if *launchEC2InstancesFlag > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := launchEC2Instances(*launchEC2InstancesFlag); err != nil {
				logger.Println("Error launching EC2 instances:", err)
			}
		}()
	}
	if *createALBFlag {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := createALB(); err != nil {
				logger.Println("Error creating ALB:", err)
			}
		}()
	}

	// Call the existing functions to create AWS resources
	wg.Add(3)
	go func() {
		defer wg.Done()
		if err := ec2.CreateEC2Instance(); err != nil {
			logger.Println("Error creating EC2 instance:", err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := elbv2.CreateLoadBalancer(); err != nil {
			logger.Println("Error creating load balancer:", err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := autoscaling.CreateAutoScalingGroup(); err != nil {
			logger.Println("Error creating auto scaling group:", err)
		}
	}()

	// Wait for all Goroutines to finish
	wg.Wait()
}

	// Call the existing functions to create AWS resources
	wg.Add(3)
	go func() {
		defer wg.Done()
		ec2.CreateEC2Instance()
	}()
	go func() {
		defer wg.Done()
		elbv2.CreateLoadBalancer()
	}()
	go func() {
		defer wg.Done()
		autoscaling.CreateAutoScalingGroup()
	}()

	// Wait for all Goroutines to finish
	wg.Wait()
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

func createEc2Instance(imageID, instanceType string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return err
	}

	// Create an EC2 instance
	ec2Client := ec2.New(sess)
	runInstancesInput := &ec2.RunInstancesInput{
		ImageId:      aws.String(imageID),
		InstanceType: aws.String(instanceType),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	_, err = ec2Client.RunInstances(runInstancesInput)
	if err != nil {
		return err
	}

	fmt.Println("EC2 instance created successfully")
	return nil
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
