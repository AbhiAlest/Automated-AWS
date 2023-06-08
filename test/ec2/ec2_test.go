package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/golang/mock/gomock"
	"github.com/AbhiAlest/Automated-AWS/config"
)

// MockEC2Client is a mocked implementation of the EC2 API
type MockEC2Client struct {
	ec2iface.EC2API
}

func (m *MockEC2Client) RunInstances(input *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	// Check if the input parameters match your test scenario and return the appropriate response or error
	if input.ImageId != nil && *input.ImageId == "ami-12345678" {
		// Return a sample response for a successful RunInstances call
		return &ec2.Reservation{
			Instances: []*ec2.Instance{
				{
					InstanceId: aws.String("i-12345678"),
				},
			},
		}, nil
	}

	// Return an error for other scenarios
	return nil, errors.New("Invalid input parameters")
}

func TestCreateEC2Instance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock EC2 client
	mockEC2 := NewMockEC2Client(ctrl)

	// Create a session using the mock EC2 client
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(config.AWSRegion),
		Credentials: config.GetAWSCredentials(),
	}))

	// Assign the mock EC2 client to the session
	svc := &ec2Service{
		client: mockEC2,
	}

	// Set up the input parameters for testing
	input := &ec2.RunInstancesInput{
		// Set the required input parameters for testing
	}

	// Set up the expected output or error for testing
	expectedOutput := &ec2.Reservation{} // Set the expected output based on your test scenario
	expectedError := nil                 // Set the expected error based on your test scenario

	// Mock the RunInstances API call on the mock EC2 client
	mockEC2.EXPECT().RunInstances(input).Return(expectedOutput, expectedError)

	// Assign the mocked session to the ec2Service
	svc.session = sess

	// Call the CreateEC2Instance function
	err := svc.CreateEC2Instance()

	// Perform assertions based on the expected output and error
	if err != expectedError {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}
	// Add more assertions based on the expected output

	// Additional test cases can be added to cover different scenarios
}

func TestCreateEC2Instance_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock EC2 client
	mockEC2 := NewMockEC2Client(ctrl)

	// Create a session using the mock EC2 client
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(config.AWSRegion),
		Credentials: config.GetAWSCredentials(),
	}))

	// Assign the mock EC2 client to the session
	svc := &ec2Service{
		client: mockEC2,
	}

	// Set up the input parameters for testing
	input := &ec2.RunInstancesInput{
		// Set the required input parameters for testing
	}

	// Set up the expected error for testing
	expectedError := awserr.New("SomeError", "Something went wrong", nil) // Set the expected error based on your test scenario

	// Mock the RunInstances API call on the mock EC2 client to return an error
	mockEC2.EXPECT().RunInstances(input).Return(nil, expectedError)

	// Assign the mocked session to the ec2Service
	svc.session = sess

	// Call the CreateEC2Instance function
	err := svc.CreateEC2Instance()

	// Perform assertions based on the expected error
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err != expectedError {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}
}

func TestCreateEC2Instance_Assertions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock EC2 client
	mockEC2 := NewMockEC2Client(ctrl)

	// Create a session using the mock EC2 client
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(config.AWSRegion),
		Credentials: config.GetAWSCredentials(),
	}))

	// Assign the mock EC2 client to the session
	svc := &ec2Service{
		client: mockEC2,
	}

	// Set up the input parameters for testing
	input := &ec2.RunInstancesInput{
		// Set the required input parameters for testing
	}

	// Set up the expected output for testing
	expectedOutput := &ec2.Reservation{
		Instances: []*ec2.Instance{
			{
				InstanceId: aws.String("i-12345678"),
			},
		},
	} // Set the expected output based on your test scenario

	// Mock the RunInstances API call on the mock EC2 client
	mockEC2.EXPECT().RunInstances(input).Return(expectedOutput, nil)

	// Assign the mocked session to the ec2Service
	svc.session = sess

	// Call the CreateEC2Instance function
	err := svc.CreateEC2Instance()

	// Perform assertions based on the expected output
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Additional assertions based on the expected output
	if len(svc.instanceIDs) != 1 {
		t.Errorf("Expected 1 instance ID, but got: %d", len(svc.instanceIDs))
	} else if svc.instanceIDs[0] != "i-12345678" {
		t.Errorf("Expected instance ID: i-12345678, but got: %s", svc.instanceIDs[0])
	}
}
