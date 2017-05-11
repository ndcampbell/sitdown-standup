package main

import (
	"github.com/nlopes/slack"
	"log"
	"os"
)

func main() {
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	api.SetDebug(true)
	log.Println("Slack Bot Starting")

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		log.Println("Event Recieved: ", msg)
	}
}
