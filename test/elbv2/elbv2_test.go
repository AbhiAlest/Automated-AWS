package elbv2_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/golang/mock/gomock"
	"github.com/AbhiAlest/Automated-AWS/config"
	"github.com/AbhiAlest/Automated-AWS/elbv2"
	mock_elbv2 "github.com/AbhiAlest/Automated-AWS/mock_elbv2"
)

func TestCreateLoadBalancer_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock ELBV2 client
	mockELBV2 := mock_elbv2.NewMockELBV2API(ctrl)

	// Create a session using the mock ELBV2 client
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(config.AWSRegion),
		Credentials: config.GetAWSCredentials(),
	}))

	// Assign the mock ELBV2 client to the ELBV2 service
	svc := &elbv2.Service{
		Client: mockELBV2,
	}

	// Set up the input parameters for testing
	input := &elbv2.CreateLoadBalancerInput{
		// Set the required input parameters for testing
	}

	// Set up the expected output for testing
	expectedOutput := &elbv2.CreateLoadBalancerOutput{
		// Set the expected output based on your test scenario
	}

	// Mock the CreateLoadBalancer API call on the mock ELBV2 client
	mockELBV2.EXPECT().CreateLoadBalancer(input).Return(expectedOutput, nil)

	// Assign the mocked session to the ELBV2 service
	svc.Session = sess

	// Call the CreateLoadBalancer function
	err := svc.CreateLoadBalancer()

	// Perform assertions based on the expected output
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Additional assertions based on the expected output
	if svc.CreatedLoadBalancerARN != expectedOutput.LoadBalancers[0].LoadBalancerArn {
		t.Errorf("Expected created load balancer ARN: %s, but got: %s",
			expectedOutput.LoadBalancers[0].LoadBalancerArn, svc.CreatedLoadBalancerARN)
	}
}

func TestCreateLoadBalancer_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock ELBV2 client
	mockELBV2 := mock_elbv2.NewMockELBV2API(ctrl)

	// Create a session using the mock ELBV2 client
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(config.AWSRegion),
		Credentials: config.GetAWSCredentials(),
	}))

	// Assign the mock ELBV2 client to the ELBV2 service
	svc := &elbv2.Service{
		Client: mockELBV2,
	}

	// Set up the input parameters for testing
	input := &elbv2.CreateLoadBalancerInput{
		// Set the required input parameters for testing
	}

	// Set up the expected error for testing
	expectedError := someError // Set the expected error based on your test scenario

	// Mock the CreateLoadBalancer API call on the mock ELBV2 client
	mockELBV2.EXPECT().CreateLoadBalancer(input).Return(nil, expectedError)

	// Assign the mocked session to the ELBV2 service
	svc.Session = sess

	// Call the CreateLoadBalancer function
	err := svc.CreateLoadBalancer()

	// Perform assertions based on the expected error
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	// Additional assertions based on the expected error
	if err != expectedError {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}
}


