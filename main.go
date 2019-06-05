package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/balabanovds/prometheus-telegram-bot/telegram"
	"github.com/balabanovds/prometheus-telegram-bot/util"
)

func main() {
	exec, err := os.Executable()
	if err != nil {
		log.Fatalf("%v", err)
	}

	util.Init(filepath.Dir(exec))

	bot, err := telegram.NewBot()
	if err != nil {
		panic(err)
	}

	bot.Run()

}
