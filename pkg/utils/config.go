package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	isSet    bool
	BotToken string `yaml:"bot_token,omitempty"`
}

var config Config

func GetConfigs() (Config, error) {
	if config.isSet {
		return config, nil
	}
	err := setConfigs()
	return config, err
}

func setConfigs() (err error) {
	configBytes, err := os.ReadFile(ConstConfigFile)
	if err != nil {
		return
	}
	if err = yaml.Unmarshal(configBytes, &config); err != nil {
		return
	}
	config.isSet = true
	return
}
