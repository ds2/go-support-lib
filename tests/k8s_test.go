package tests

import (
	"flag"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ds_2/go-support-lib/k8s"
)

var k8sUrl = flag.String("k8sBaseUrl", "http://localhost:6443", "defines the k8s base url to access for testing")
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
	nodes := k8s.GetNodesInternalIpAddresses(*k8sUrl, *nodeCount, *k8sAccessToken, "")

	if nodes == nil || len(nodes) <= 0 {
		t.Error("No nodes are returned!")
	}
}

func TestAwsEksNodeList(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping GetNodesInternalIpAddresses Test")
	}
	nodes := k8s.GetNodesInternalIpAddresses(*k8sUrl, *nodeCount, k8s.GetAwsK8sAccessToken(*awsK8sClusterId), "")
	if nodes == nil || len(nodes) <= 0 {
		t.Error("No nodes are returned!")
	}
}

func TestGetAwsIamToken(t *testing.T){
	token:=k8s.GetTokenViaAwsIamAuthenticatorClient(*awsK8sClusterId)
	assert.NotNil(t, token, "Token is null")
	assert.NotEmpty(t, token, "Token is empty")
}