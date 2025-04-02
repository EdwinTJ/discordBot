package main

import (
	"discordbot/bot"
	"discordbot/config"
	"fmt"
)

func main(){
	err := config.ReadConfig()

	if err !=nil{
		fmt.Println("Error reading config:", err)
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}