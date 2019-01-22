package k8s

import (
	"fmt"
	"log"

	"os/exec"

	"github.com/kubernetes-sigs/aws-iam-authenticator/pkg/token"
	"github.com/sirupsen/logrus"
)

func GetAwsK8sAccessToken(k8sClusterId string) (accessToken string) {
	var verifierSvc = token.NewVerifier(k8sClusterId)
	thisGen, _ := token.NewGenerator()
	awsToken, _ := thisGen.Get(k8sClusterId)
	accessToken = awsToken
	logrus.Debugln("token should be ", accessToken)
	var identity, err = verifierSvc.Verify(accessToken)
	if err != nil {
		logrus.Error("Issue: ", err)
	}
	logrus.Debugln("Identity found: ", identity)
	return accessToken
}

func GetTokenViaAwsIamAuthenticatorClient(k8sClusterId string) (accesstoken string) {
	//aws-iam-authenticator
	out, err := exec.Command("aws-iam-authenticator", "token", "-i", k8sClusterId).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
	return accesstoken
}
