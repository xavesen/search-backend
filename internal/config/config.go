package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	KafkaAddrsStr			string		`mapstructure:"KAFKA_ADDR"`
	KafkaAddrs 				[]string
	KafkaTopic				string		`mapstructure:"KAFKA_TOPIC"`
	KafkaGroupId			string		`mapstructure:"KAFKA_GROUP_ID"`
	KafkaCommitInterval		int			`mapstructure:"KAFKA_COMMIT_INTERVAL"`

	ElasticSearchURLsStr	string		`mapstructure:"ELASTIC_SEARCH_URL"`
	ElasticSearchURLs		[]string
	ElasticSearchKey		string		`mapstructure:"ELASTIC_SEARCH_KEY"`

	LogLevel 				log.Level	`mapstructure:"LOG_LEVEL"`
}

func LoadConfig() (*Config, error) {
	log.Info("Loading config from environment")
	var config Config

	viper.AutomaticEnv()
	log.Info("Parsing environment variables to config struct")
	if err := viper.Unmarshal(&config); err != nil {
		log.Errorf("Error parsing environment variables to config struct: %s", err.Error())
		return nil, err
	}

	config.KafkaAddrs = strings.Split(config.KafkaAddrsStr, ";")
	config.ElasticSearchURLs = strings.Split(config.ElasticSearchURLsStr, ";")

	log.Infof("Setting log level to %s", config.LogLevel.String())
	log.SetLevel(config.LogLevel)

	log.Info("Successfully loaded config from environment")
	return &config, nil
}