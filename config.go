package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type config struct {
	CredsFile           string `json:"credsFile"`
	AppID               string `json:"appID"`
	SlackURL            string `json:"slackURL"`
	OnlyNotifyOnNotPass bool   `json:"onlyNotifyOnNotPass"`
	IncludePolicyStatus bool   `json:"includePolicyStatus"`
}

func parseConfig(configFile string) config {
	var config config

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
