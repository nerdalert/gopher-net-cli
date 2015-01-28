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

var GnetCtlShow = cli.Command{
	Name:  "show",
	Usage: "gnet-ctl show <option>",
	Subcommands: []cli.Command{
		{
			Usage:  "use 'gnet-ctl show help' for subcommand usage",
			Action: cli.ShowSubcommandHelp,
		},
		{
			Name:   "neighbors-config",
			Usage:  "'gnet-ctl show neighbors-config' - show neighbor configuration and current state",
			Action: ShowNeighborsConfigs,
		},
		{
			Name:   "routes",
			Usage:  "'gnet-ctl show routes' - show best incoming destinations",
			Action: ShowGetRoutes,
		},
		{
			Name:   "neighbors",
			Usage:  "'gnet-ctl show neighbors' - show neighbor configurations",
			Action: ShowNeighbors,
		},
		{
			Name:   "global-config",
			Usage:  "'gnet-ctl show global-config' - show the local BGP global configuration",
			Action: ShowGlobalConfig,
		},
		{
			Name:   "rib-out",
			Usage:  "'gnet-ctl show rib-out' - show the bgp rib-out table",
			Action: ShowRibOut,
		},
		{
			Name:   "rib-in",
			Usage:  "'gnet-ctl show rib-in' - show the bgp rib-in table",
			Action: ShowRibIn,
		},
	},
}

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

func ShowNeighborsConfigs(c *cli.Context) {
	log.Debugln("get neighbor configs CLI called")
	client := NewClient()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080"+NEIGHBORS_CONFIG, body)
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

func ShowGlobalConfig(c *cli.Context) {
	log.Debugln("get routes CLI called")
	client := NewClient()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080"+GLOBAL_CONFIG, body)
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


func ShowGetRoutes(c *cli.Context) {
    log.Debugln("get routes CLI called")
    client := NewClient()
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    writer.Close()
    req, err := http.NewRequest("GET", "http://127.0.0.1:8080"+ROUTE_TABLES, body)
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

func ShowRibIn(c *cli.Context) {
    log.Debugln("get RIB_IN CLI called")
    client := NewClient()
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    writer.Close()
    req, err := http.NewRequest("GET", "http://127.0.0.1:8080"+RIB_IN, body)
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

func ShowRibOut(c *cli.Context) {
    log.Debugln("get RIB_OUT CLI called")
    client := NewClient()
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    writer.Close()
    req, err := http.NewRequest("GET", "http://127.0.0.1:8080"+RIB_OUT, body)
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
