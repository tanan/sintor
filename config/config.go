package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Application AppConfig `yaml:"application"`
}

type AppConfig struct {
	Port     int          `toml:"port"`
	Database DBConfig     `yaml:"database"`
	Notify   NotifyConfig `yaml:"notify"`
}

type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	Host     string `yaml:"host"`
	Schema   string `yaml:"schema"`
	Port     int    `yaml:"port"`
}

type NotifyConfig struct {
	Slack SlackConfig `yaml:"slack"`
}

type SlackConfig struct {
	Token   string `yaml:"token"`
	ApiURL  string `yaml:"api_url"`
	Channel string `yaml:"channel"`
	User    string `yaml:"user"`
}

func LoadFile(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
