package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
	"strings"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Println("SLACK_TOKEN not found")
		os.Exit(1)
	}
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	//api.SetDebug(true)
	log.Println("Slack Bot Starting")

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			info := rtm.GetInfo()
			if ev.User != info.User.ID { //verifies the message is not from the bot
				if info.GetIMByID(ev.Channel).IsIM {
					log.Printf("Direct message: %s", ev.Text)
					go respond(rtm, ev)
				} else {
					log.Printf("Message: %v %v", ev.Text, ev.User)
				}
			}
		case *slack.RTMError:
			log.Printf("Error: %s\n", ev.Error())
		case *slack.InvalidAuthEvent:
			log.Printf("Invalid credentials")
			os.Exit(1)
		default:
			// ignore other events
		}
	}
}

//function to handle direct messages to bot
func respond(rtm *slack.RTM, msg *slack.MessageEvent) {
	var response string
	args := strings.Split(msg.Text, " ")
	switch args[0] {
	case "help":
		help := "Supported commands:\n"
		response = fmt.Sprintf(help)
	}
	rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
}
