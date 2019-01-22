package tests

import (
	"flag"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ds_2/go-support-lib/aws"
	"gitlab.com/ds_2/go-support-lib/common"
)

var elbName = flag.String("testElbName", "", "the test elb name to use for testing")

func TestLbHealth(t *testing.T) {
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	elbSvc := elbv2.New(awsSession)
	state := aws.GetElbState(elbSvc, *elbName)
	assert.NotNil(t, state, "State is null")
	assert.NotEqual(t, common.State_Unknown, state, "State should not be unknown for a known LB")
}

func TestLbArn(t *testing.T) {
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	elbSvc := elbv2.New(awsSession)
	state := aws.GetLoadBalancerArn(elbSvc, *elbName)
	assert.NotNil(t, state, "ARN is null")
	assert.NotEqual(t, "", state, "ARN not set")
}

func TestTargetHealth(t *testing.T) {
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	elbSvc := elbv2.New(awsSession)
	state := aws.GetElbTargetHealth(elbSvc, *elbName)
	assert.NotNil(t, state, "State is null")
	assert.NotEqual(t, common.State_Unknown, state, "State is unknown from the given target!")
}
