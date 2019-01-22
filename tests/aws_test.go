package tests

import (
	"flag"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ds_2/go-support-lib/aws"
	"gitlab.com/ds_2/go-support-lib/common"
)

var elbName = flag.String("testElbName", "", "the test elb name to use for testing")

func TestLbHealth(t *testing.T) {
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	state := aws.GetElbState(awsSession, *elbName)
	assert.NotNil(t, state, "State is null")
	assert.NotEqual(t, common.Unknown, state, "State should not be unknown for a known LB")
}
