package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	log "github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func ShowNeighbors(c *cli.Context) {
	log.Debugln("get routes CLI called")
	client := NewClient()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080"+NEIGHBORS, body)
	if err != nil {
		log.Error(error(err))
		return
	}
	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		log.Error(error(parseFormErr))
		return
	}
	// Get Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("No answer from the bgp daemon: %s \n", err)
		return
	}

	// Read Response
	respBody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return
	}
	if resp.Status == "404 Not Found" {
		log.Println("requested data type not supported")
	}
	// Display Results
	log.Debugln("response Status : ", resp.Status)
	log.Debugln("response Headers : ", resp.Header)
	log.Debugln("response Body : ", string(respBody))
	log.Infoln("results: ", string(respBody))
	return
}
