package tests

import (
	"flag"
	"testing"

	"github.com/sirupsen/logrus"
	"gitlab.com/ds_2/go-support-lib/k8s"
)

var awsUrl = flag.String("k8sBaseUrl", "", "defines the k8s base url to access for testing")
var k8sAccessToken = flag.String("k8sToken", "", "The access token to use")
var awsK8sClusterId = flag.String("awsEksClusterId", "", "The AWS EKS clusterId")
var nodeCount = flag.Uint("nodeCount", 10, "the node count to return")

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func TestGetNodeList(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping GetNodesInternalIpAddresses Test")
	}
	nodes := k8s.GetNodesInternalIpAddresses(*awsUrl, *nodeCount, *k8sAccessToken, "")

	if nodes == nil || len(nodes) <= 0 {
		t.Error("No nodes are returned!")
	}
}

func TestAwsEksNodeList(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping GetNodesInternalIpAddresses Test")
	}
	nodes := k8s.GetNodesInternalIpAddresses(*awsUrl, *nodeCount, k8s.GetAwsK8sAccessToken(*awsK8sClusterId), "")
	if nodes == nil || len(nodes) <= 0 {
		t.Error("No nodes are returned!")
	}
}
