package tests

import (
	"testing"

	"gitlab.com/ds_2/go-support-lib/k8s"
)

func TestGetNodeList(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping GetNodeList Test")
	}

	nodes := k8s.GetNodeList("https://DE103F2619FFEAF75C518C4C6F620ECA.yl4.eu-central-1.eks.amazonaws.com", 10)

	if nodes == nil {
		t.Error("No nodes are returned!")
	}
}
