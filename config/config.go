package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	WebHook struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Token string `yaml:"token"`
		// Port is the local machine TCP Port to bind the HTTP Server to
		Url    string `yaml:"url"`
		Port   string `yaml:"port"`
		ChatID string `yaml:"chatId"`
	} `yaml:"webhook"`
}

func Init(configPath string) *Config {
	// Create config structure
	c := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	//defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&c); err != nil {
		panic(err)
	}

	return c
}
