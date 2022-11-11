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
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// Returns the internal ip addresses of all known nodes of the cluster.
func GetNodesInternalIpAddresses(clusterEndpoint string, limit uint, authToken string, certClientFile string) (result []string) {
	var nodeUrl = clusterEndpoint + "/api/v1/nodes?limit=" + strconv.Itoa(int(limit)) + "&pretty=true"
	logrus.Debug("URL to use is: ", nodeUrl)
	var cert tls.Certificate
	if len(certClientFile) > 0 {
		logrus.Debug("Adding client cert from ", certClientFile)
		cert, _ = tls.LoadX509KeyPair(certClientFile, "")
	}
	var tlsConfig = tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}
	if len(certClientFile) <= 0 {
		tlsConfig = tls.Config{
			InsecureSkipVerify: true,
		}
	}
	tr := &http.Transport{
		TLSClientConfig: &tlsConfig,
	}
	var httpClient = &http.Client{Timeout: 9 * time.Second, Transport: tr}
	var req, err = http.NewRequest("GET", nodeUrl, nil)
	if err != nil {
		logrus.Fatal("Could not setup http client! ", err)
	}
	if len(authToken) > 0 {
		logrus.Debug("Adding bearer token..", authToken)
		req.Header.Set("Authorization", "Bearer "+authToken)
	}
	var body, err2 = httpClient.Do(req)
	if err2 != nil {
		logrus.Error("Error when accessing the service: ", err2)
		return nil
	}
	defer body.Body.Close()
	logrus.Debug("Return status: ", body.StatusCode)
	logrus.Debug("Return length: ", body.ContentLength)
	var nodeListData *NodeList
	var bodyBytes, _ = ioutil.ReadAll(body.Body)
	ioutil.WriteFile(".awsResponse.json", bodyBytes, 0644)
	bodyString := string(bodyBytes)
	if body.StatusCode >= 400 {
		logrus.Error("No result data found!")
		return nil
	}
	logrus.Debug("Response: ", bodyString)
	err = json.Unmarshal(bodyBytes, nodeListData)
	if err != nil {
		logrus.Error("Could not parse response body into NodeList type! ", err)
		return nil
	}
	logrus.Debug("NodeLIst: ", nodeListData)
	for _, nodeListItem := range nodeListData.Items {
		logrus.Debug("NodeListItem: ", nodeListItem)
		for _, addr := range nodeListItem.Status.Addresses {
			logrus.Debug("Address: ", addr)
			var addrType = NodeAddressType_value[addr.Type]
			if addrType < 0 {
				continue
			}
			var typeResolved NodeAddressType
			typeResolved = NodeAddressType(addrType)
			if NodeAddressType_InternalIP == typeResolved {
				result = append(result, addr.Address)
			}
		}
	}
	logrus.Debug("Result: ", result)
	return result
}

//func homeDir() string {
//	if h := os.Getenv("HOME"); h != "" {
//		return h
//	}
//	return os.Getenv("USERPROFILE") // windows
//}
//
//func rolloutNewImageToDeployment(details RolloutImageDetails) {
//	var kubeconfig *string
//	if home := homeDir(); home != "" {
//		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
//	} else {
//		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
//	}
//	flag.Parse()
//
//	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// create the clientset
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		panic(err.Error())
//	}
//	deploymentData, err:=clientset.AppsV1().Deployments(details.Namespace).Get(details.DeploymentName, metav1.GetOptions{})
//	if err != nil {
//		panic(err.Error())
//	}
//
//
//}
