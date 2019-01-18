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

func GetNodeList(clusterEndpoint string, limit int, authToken string, certClientFile string) (result []string) {
	var nodeUrl = clusterEndpoint + "/api/v1/nodes?limit=" + strconv.Itoa(limit) + "&pretty=true"
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
		logrus.Debug("Adding bearer token..")
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
	var nodeListData NodeList
	var bodyBytes, _ = ioutil.ReadAll(body.Body)
	ioutil.WriteFile("awsResponse.json", bodyBytes, 0644)
	bodyString := string(bodyBytes)
	logrus.Debug("Response: ", bodyString)
	err = json.Unmarshal(bodyBytes, &nodeListData)
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
	return result
}
