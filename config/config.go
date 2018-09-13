package config

import "github.com/BurntSushi/toml"

type Config struct {
	Application AppConfig   `toml:"application"`
	Slack       SlackConfig `toml:"slack"`
}

type AppConfig struct {
	Port int `toml:"port"`
}

type SlackConfig struct {
	Token   string `toml:"token"`
	ApiURL  string `toml:"api_url"`
	Channel string `toml:"channel"`
	User    string `toml:"user"`
}

func LoadFile(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
