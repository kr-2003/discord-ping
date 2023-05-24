package main

import (
	"app/main/config"
	"app/main/bot"
	"fmt"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
