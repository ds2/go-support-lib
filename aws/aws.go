package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/sirupsen/logrus"
	"gitlab.com/ds_2/go-support-lib/common"
)

func GetElbState(session *session.Session, elbName string) (state common.State) {
	state = common.Unknown
	elbSvc := elbv2.New(session)
	foundLoadBalancers, err := elbSvc.DescribeLoadBalancers(&elbv2.DescribeLoadBalancersInput{
		Names: []*string{&elbName},
	})
	if err != nil {
		logrus.Error("Error when getting the load balancers: ", err.Error())
	}
	var countLoadBalancers = len(foundLoadBalancers.LoadBalancers)
	if (countLoadBalancers <= 0) {
		logrus.Warn("No load balancers found by name of ", elbName)
	} else {
		//there are some load balancers ;)
		for _, thisLb := range foundLoadBalancers.LoadBalancers {
			lbArn := *thisLb.LoadBalancerArn
			lbState := *thisLb.State.Code
			logrus.Debugln("ARN is ", lbArn, ", state is ", lbState)
			switch lbState {
			case "active":
				state = common.Active
			default:
				state = common.Unknown
			}
		}
	}
	logrus.Debugln("Returning state: ", state)
	return state
}
