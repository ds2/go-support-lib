package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gitlab.com/ds_2/go-support-lib/sysinfo"
)

func SendAsJson(url string, t *sysinfo.HealthInfo) {
	log.Println("Dto values: ", t)
	var jsonBytes, err = json.Marshal(t)
	if err != nil {
		panic(err)
	}
	log.Println("Json: ", jsonBytes)
	//enc := json.NewEncoder(os.Stdout)
	//enc.Encode(t)
	//var nodeName = "myNodeUuid"
	//url := "https://n8w8.ds-2.de/api/rs/v1/health/" + nodeName
	fmt.Println("URL:>", url)

	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
