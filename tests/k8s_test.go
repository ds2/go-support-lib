package tests

import (
	"flag"
	"testing"

	"github.com/sirupsen/logrus"
	"gitlab.com/ds_2/go-support-lib/k8s"
)

var awsUrl = flag.String("k8sBaseUrl", "", "defines the k8s base url to access for testing")
var k8sAccessToken = flag.String("k8sToken", "", "The access token to use")

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func TestGetNodeList(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping GetNodeList Test")
	}
	nodes := k8s.GetNodeList(*awsUrl, 10, *k8sAccessToken, "")

	if nodes == nil || len(nodes) <= 0 {
		t.Error("No nodes are returned!")
	}
}
