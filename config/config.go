package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var(
	Token string
	BotPrefix string
	ChannelID string

	config *configStruct
)

type configStruct struct {
	Token string `json:"Token"`
	BotPrefix string `json:"Bot_prefix"`
	ChannelID string `json:"Channel_id"`
}

func ReadConfig() error {
	fmt.Println("Reading config.json file...")
	file,err := ioutil.ReadFile("config.json")

	if err != nil {
		return fmt.Errorf("Error reading config.json file: %v", err)

	}

	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)
	if err != nil {
		return fmt.Errorf("Error unmarshalling config.json file: %v", err)
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	ChannelID = config.ChannelID


	return nil
}

