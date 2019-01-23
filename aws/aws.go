package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/sirupsen/logrus"
	"gitlab.com/ds_2/go-support-lib/common"
)

func GetElbState(elbSvc *elbv2.ELBV2, elbName string) (state common.State) {
	state = common.State_Unknown
	foundLoadBalancers, err := elbSvc.DescribeLoadBalancers(&elbv2.DescribeLoadBalancersInput{
		Names: []*string{&elbName},
	})
	if err != nil {
		logrus.Error("Error when getting the load balancers: ", err.Error())
	}
	var countLoadBalancers = len(foundLoadBalancers.LoadBalancers)
	if countLoadBalancers <= 0 {
		logrus.Warn("No load balancers found by name of ", elbName)
	} else {
		//there are some load balancers ;)
		for _, thisLb := range foundLoadBalancers.LoadBalancers {
			lbArn := *thisLb.LoadBalancerArn
			lbState := *thisLb.State.Code
			logrus.Debugln("ARN is ", lbArn, ", state is ", lbState)
			switch lbState {
			case "active":
				state = common.State_Active
			default:
				logrus.Warn("Unmapped lb state: ", lbState)
			}
		}
	}
	logrus.Debugln("Returning state: ", state)
	return state
}

// Handles an AWS error somehow. Returns true if there was an error, otherwise false.
func handleAwsError(error2 error, errorMsg string, fail bool) bool {
	if error2 != nil {
		if fail {
			logrus.Fatalln(errorMsg, error2.Error())
		} else {
			logrus.Warn(errorMsg, error2.Error())
		}
		return true
	}
	return false
}

//Returns the ARNs for the target groups of this load balancer.
func GetElbTargetGroupsArns(elbSvc *elbv2.ELBV2, elbName string) (tgArns []string) {
	lbArn := GetLoadBalancerArn(elbSvc, elbName)
	targetHealthResponse, err := elbSvc.DescribeTargetGroups(&elbv2.DescribeTargetGroupsInput{
		LoadBalancerArn: &lbArn,
	})
	if !handleAwsError(err, "Error when getting the targetGroups for the given LB", false) {
		for _, tg := range targetHealthResponse.TargetGroups {
			tgArns = append(tgArns, *tg.TargetGroupArn)
		}
	}
	logrus.Debugln("targetGroup ARNs are: ", tgArns)
	return tgArns
}

func GetElbInstanceHealthViaSession(session *session.Session, elbName string, instanceIds []string) common.State {
	return GetElbInstanceHealth(elbv2.New(session), elbName, instanceIds)
}

//Returns the common health state of the targets of the given load balancer.
func GetElbInstanceHealth(elbSvc *elbv2.ELBV2, elbName string, instanceIds []string) (state common.State) {
	state = common.State_Unknown
	tgArn := GetElbTargetGroupsArns(elbSvc, elbName)
	if len(tgArn) <= 0 {
		logrus.Errorln("No targetGroups found for ELB ", elbName)
	}
	var targetDescriptions []*elbv2.TargetDescription
	for _, instId := range instanceIds {
		targetDescriptions = append(targetDescriptions, &elbv2.TargetDescription{
			Id: &instId,
		})
	}
	targetHealthResponse, err := elbSvc.DescribeTargetHealth(&elbv2.DescribeTargetHealthInput{
		TargetGroupArn: &tgArn[0],
		Targets:        targetDescriptions,
	})
	if !handleAwsError(err, "Error when trying to get the target health", false) {
		countTargets := len(targetHealthResponse.TargetHealthDescriptions)
		if countTargets <= 0 {
			logrus.Warn("No targets found in ", tgArn)
		} else {
			for _, tg := range targetHealthResponse.TargetHealthDescriptions {
				logrus.Debugln("This target: ", tg)
				stateStr := *tg.TargetHealth.State
				logrus.Debugln("State of this target: ", stateStr)
				switch stateStr {
				case "active", "healthy":
					state = common.State_Active
				case "unhealthy":
					state = common.State_Error
					break
				case "initial":
					state=common.State_New
					break
				default:
					logrus.Warn("Unmapped state: ", stateStr)
				}
			}
		}
	}
	return state
}

// Returns the ARN of the given load balancer.
func GetLoadBalancerArn(elbSvc *elbv2.ELBV2, elbName string) (arn string) {
	arn = ""
	foundLoadBalancers, err := elbSvc.DescribeLoadBalancers(&elbv2.DescribeLoadBalancersInput{
		Names: []*string{&elbName},
	})
	if err != nil {
		logrus.Error("Error when getting the load balancers: ", err.Error())
	} else {
		var countLoadBalancers = len(foundLoadBalancers.LoadBalancers)
		if countLoadBalancers <= 0 {
			logrus.Warn("No load balancers found by name of ", elbName)
		} else {
			//there are some load balancers ;)
			for _, thisLb := range foundLoadBalancers.LoadBalancers {
				arn = *thisLb.LoadBalancerArn
				break
			}
		}
	}
	logrus.Debugln("ARN found is ", arn)
	return arn
}
