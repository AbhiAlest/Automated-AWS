#!/bin/bash

# Set environment variables
export AWS_REGION="us-west-2"
export AWS_PROFILE="your-aws-profile"

# Build and run the test
go test -v ./autoscaling_test.go

# Unset environment variables
unset AWS_REGION
unset AWS_PROFILE
