//go:generate mockgen -destination=elbv2_mock.go -package=mocks github.com/aws/aws-sdk-go/service/elbv2 ELBV2API
package mocks
import "github.com/aws/aws-sdk-go/service/elbv2"
type ELBV2API interface {
   // Declare the methods you want to mock from the elbv2.ELBV2API interface
   CreateLoadBalancer(*elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error)
   // ...
}
