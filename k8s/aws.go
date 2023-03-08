// Copyright 2022 Dirk Strauss
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8s

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/sirupsen/logrus"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

func GetAwsK8sAccessToken(k8sClusterId string, stsPartitionId string, awsRegion string) (accessToken string) {
	var verifierSvc = token.NewVerifier(k8sClusterId, stsPartitionId, awsRegion)
	thisGen, _ := token.NewGenerator(true, true)
	awsToken, _ := thisGen.Get(k8sClusterId)
	accessToken = awsToken.Token
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
	var answer *AwsIamAuthenticatorResponse
	err = json.Unmarshal(out, answer)
	if err != nil {
		log.Fatalln("Error when trying to unmarshal the received data from the aws iam authenticator: ", err)
	}
	logrus.Debugln("Answer looks like: ", answer)
	accesstoken = answer.Status.Token
	logrus.Debugln("The accessToken might be: ", accesstoken)
	return accesstoken
}
