package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/balabanovds/prometheus-telegram-bot/telegram"
	"github.com/balabanovds/prometheus-telegram-bot/util"
)

func main() {
	util.Init()

	bot, err := telegram.NewBot()
	if err != nil {
		panic(err)
	}

	bot.Run()

}
