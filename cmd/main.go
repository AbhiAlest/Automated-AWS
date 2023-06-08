package main

import (
	"github.com/AbhiAlest/Automated-AWS/pkg/ec2"
	"github.com/AbhiAlest/Automated-AWS/pkg/elbv2"
	"github.com/AbhiAlest/Automated-AWS/pkg/autoscaling"
)

func main() {
	// Call the functions to create AWS resources
	ec2.CreateEC2Instance()
	elbv2.CreateLoadBalancer()
	autoscaling.CreateAutoScalingGroup()
}

