package main

import (
	"encoding/json"
	"net/http"
)

// Config はシステム設定です
type Config struct {
	Salt   string `json:"salt"`
	Strech int    `json:"strech"`
}

var conf *Config

func getConfig(r *http.Request) (*Config, error) {
	if conf != nil {
		return conf, nil
	}
	// Firestore から設定読み込み
	confstr, err := ds.Get(r, confPRE+"Config")
	if err == nil {
		var c Config
		if err := json.Unmarshal([]byte(confstr), &c); err == nil {
			conf = &c
			return conf, nil
		}
	}
	return nil, err
}
