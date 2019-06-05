package util

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config main
type config struct {
	Telegram struct {
		Token   string  `yaml:"token"`
		Cert    string  `yaml:"cert"`
		Key     string  `yaml:"key"`
		Webhook string  `yaml:"webhook-url"`
		Chats   []int64 `yaml:"chats"`
	} `yaml:"telegram"`

	Tor struct {
		HTTPProxy  string `yaml:"http-proxy"`
		SOCKSProxy string `yaml:"socks-proxy"`
	} `yaml:"tor"`

	Prometheus struct {
		API     string  `yaml:"api-url"`
		Queries []Query `yaml:"qeuries"`
	} `yaml:"prometheus"`

	Service struct {
		HTTPAddr string `yaml:"http-addr"`
		TLSAddr  string `yaml:"tls-addr"`
	} `yaml:"service"`
}

// Query basic struct
type Query struct {
	Name  string `yaml:"name"`
	Query string `yaml:"query"`
}

var (
	// Cfg is global config variable
	Cfg config
)

// Init reads config
func Init() {
	exec, err := os.Executable()
	if err != nil {
		log.Fatalf("%v", err)
	}
	baseDir := filepath.Dir(exec)
	log.Println(baseDir)
	// data, err := ioutil.ReadFile(filepath.Join(baseDir, "config.yml"))
	data, err := ioutil.ReadFile("/home/dbalaban/projects/pointer/prometheus-telegram-bot/config.yml")
	if err != nil {
		log.Fatalf("ReadFile() config err %v", err)
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		log.Fatalf("Unmarshal() config err %v", err)
	}
}
