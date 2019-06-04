package main

import (
	"log"
	"net/http"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	proxyURL, err := url.Parse("http://localhost:8118")
	if err != nil {
		log.Panic(err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Transport: transport,
	}

	bot, err := tgbotapi.NewBotAPIWithClient("888291143:AAHN1aw2jFZWFRKwPcZ2D9fTHe28cQcBh08", client)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
}
