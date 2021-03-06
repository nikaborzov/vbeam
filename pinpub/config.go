package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	IpfsRepoPath string `json:"ipfsRepoPath"`
	Topics       map[string]struct {
		SrcDir string `json:"srcDir"`
	} `json:"topics"`
}

func loadConfig() (*config, error) {
	byt, err := ioutil.ReadFile("./pinpub.json")
	if err != nil {
		return nil, err
	}

	cfg := config{}
	err = json.Unmarshal(byt, &cfg)
	if err != nil {
		return nil, err
	}
	fmt.Printf("config is now %#v\n", cfg)
	return &cfg, nil
}
