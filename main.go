package main

import (
	tgClient "awesomeProject3/clients/telegram"
	event_consumer "awesomeProject3/consumer/event-consumer"
	"awesomeProject3/events/telegram"
	"awesomeProject3/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

//read_adviser_jewps_bot

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath))

	log.Print("Starting bot")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal(err)
	}
}

func mustToken() string {
	// bot -tg-bot-token 'token'
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
