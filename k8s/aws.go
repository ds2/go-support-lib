package k8s

import (
	"encoding/json"
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
	logrus.Debugln("Trying to get token for cluster ", k8sClusterId)
	out, err := exec.Command("aws-iam-authenticator", "token", "-i", k8sClusterId).Output()
	if err != nil {
		log.Fatal("Error occurred: ", err, ": ", out)
	}
	//response is a json string
	var answer AwsIamAuthenticatorResponse
	err = json.Unmarshal(out, &answer)
	if err != nil {
		log.Fatalln("Error when trying to unmarshal the received data from the aws iam authenticator: ", err)
	}
	logrus.Debugln("Answer looks like: ", answer)
	accesstoken = answer.Status.Token
	logrus.Debugln("The accessToken might be: ", accesstoken)
	return accesstoken
}
