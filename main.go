package main

import (
	tgClient "awesomeProject3/clients/telegram"
	event_consumer "awesomeProject3/consumer/event-consumer"
	"awesomeProject3/events/telegram"
	"awesomeProject3/storage/sqlite"
	"context"
	"flag"
	"log"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

//read_adviser_jewps_bot

func main() {

	//s := files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatalf("Cant connect to storage: %v", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("cant init storage:", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("Starting bot")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal(err)
	}
}

func mustToken() string {
	// bot -tg-bot-token 'token'
	token := flag.String(
		"token",
		"",
		"token for access to bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
