package main

import (
	"./config"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
)

const (
	// TOKEN telegram
	TOKEN = "5536309016:AAFZgNln1RjUwuhNAQlv2FB2vy24Wqmsv8c"
	// URL telegram
	URL = "https://api.telegram.org/bot"
	// PORT local
	PORT = "3000"
)

func main() {
	cfg, err := NewConfig("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	//http.HandleFunc("/api/v1/update", update)
	fmt.Println("Listening on port", cfg.WebHook.Port, ".")
	if err := http.ListenAndServe(":"+cfg.WebHook.Port, nil); err != nil {
		log.Fatal(err)
	}
}

func NewConfig(configPath string) (*config.Config, error) {
	// Create config structure
	c := &config.Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	//defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&c); err != nil {
		return nil, err
	}

	return c, nil
}
