package k8s

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func getNodeList(clusterEndpoint string) (result []string) {
	var nodeUrl = clusterEndpoint + "/api/v1/nodes?limit=10"
	var httpClient = &http.Client{Timeout: 9 * time.Second}
	var body, err = httpClient.Get(nodeUrl)
	if (err != nil) {
		logrus.Error("Error when accessing the service: ", err)
		return nil
	}
	defer body.Body.Close()
	var nodeListData NodeList
	err = json.NewDecoder(body.Body).Decode(nodeListData)
	if (err != nil) {
		logrus.Error("Could not parse response body into NodeList type!")
		return nil
	}
	for _, nodeListItem := range nodeListData.Items {
		logrus.Debug("NodeListItem: ", nodeListItem)
	}
	return result
}
