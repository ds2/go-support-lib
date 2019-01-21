package k8s

import "github.com/sirupsen/logrus"
import "github.com/kubernetes-sigs/aws-iam-authenticator/pkg/token"

func GetAwsK8sAccessToken(k8sClusterId string) (accessToken string) {
	thisGen, _ := token.NewGenerator(true)
	awsToken, _ := thisGen.Get(k8sClusterId)
	accessToken = awsToken.Token
	logrus.Debugln("token should be ", accessToken)
	return accessToken
}
