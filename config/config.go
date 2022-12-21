package config

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
