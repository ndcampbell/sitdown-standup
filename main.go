package main

import (
	"github.com/nlopes/slack"
	"log"
	"os"
)

func main() {
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	//api.SetDebug(true)
	log.Println("Slack Bot Starting")

	rtm := api.NewRTM()
	go rtm.ManageConnection()
	info := rtm.GetInfo()
	log.Printf("Bot User: %v", info.User.Name)
	botID := info.User.ID

	for msg := range rtm.IncomingEvents {
		log.Println("Event Recieved: ")
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			log.Printf("Message: %v %v", ev.Text, ev.User)
		default:
			// ignore other events
		}
	}
}
