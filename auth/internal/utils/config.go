package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	PORT           string `json:"port,omitempty"`
	JWT_ISSUER     string `json:"jwt_issuer,omitempty"`
	JWT_EXPIRATION string `json:"jwt_expiration,omitempty"`
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
