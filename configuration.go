package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func registerAgent(endPoint, name string, key string) ([]byte, error) {
	var err error
	payload := strings.NewReader(fmt.Sprintf("{\n\"agentName\": \"%s\"\n}", name))
	if req, err := http.NewRequest("POST", endPoint, payload); err == nil {
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("UTM-Token", key)
		if res, err := http.DefaultClient.Do(req); err == nil {
			defer res.Body.Close()
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				return body, nil
			}
		}
	}
	return nil, err
}

type config struct {
	Server   string `yaml:"server"`
	AgentID  string `yaml:"agent-id"`
	AgentKey string `yaml:"agent-key"`
}

var oneConfigRead sync.Once
var cnf config

func readConfig() {
	err := readYAML("config.yml", &cnf)
	if err != nil {
		h.FatalError("error reading config %v", err)
	}
}

func getConfig() config {
	oneConfigRead.Do(func() { readConfig() })
	return cnf
}

func writeConfig(cnf config) error {
	err := writeYAML("config.yml", cnf)
	if err != nil {
		return err
	}
	return nil
}
