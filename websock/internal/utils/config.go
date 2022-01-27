package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	PORT               string `json:"port,omitempty"`
	JWT_ISSUER         string `json:"jwt_issuer,omitempty"`
	JWT_EXPIRATION     string `json:"jwt_expiration,omitempty"`
	KAFKA_SERVICE      string `json:"kafka_service,omitempty"`
	REDIS_SERVICE      string `json:"redis_service,omitempty"`
	KAFKA_PORT         string `json:"kafka_port,omitempty"`
	REDIS_PORT         string `json:"redis_port,omitempty"`
	WORKER_COUNT       string `json:"worker_count,omitempty"`
	USER_TABLE_NAME    string `json:"user_table_name,omitempty"`
	MESSAGE_TABLE_NAME string `json:"message_table_name,omitempty"`
	AWS_REGION         string `json:"aws_region,omitempty"`
}

func LoadConfig(filename string) (Config, error) {
	var conf Config

	confFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return Config{}, errors.New("unable to red config file")
	}
	defer confFile.Close()
	jsonParser := json.NewDecoder(confFile)
	jsonParser.Decode(&conf)
	return conf, nil
}
